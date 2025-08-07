package service

import (
	"api/internal/consts"
	"api/internal/dao"
	"api/internal/model"
	"api/internal/model/entity"
	"api/utility"
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/filecoin-project/lotus/api/client"
	"github.com/filecoin-project/lotus/api/v2api"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
)

var Staking = stakingService{}

type stakingService struct {
	Api v2api.FullNode
	Err error
}

func (s *stakingService) Start(ctx context.Context) {
	c := s.Init(ctx)
	if c.Err != nil {
		g.Log().Error(ctx, "init staking service error", c.Err)
		return
	}
	pending := []entity.SliverPlan{}
	err := dao.SliverPlan.Ctx(ctx).Where("status", consts.ConstPlanStatusPending).WithAll().Scan(&pending)
	if err != nil {
		g.Log().Error(ctx, "get pending proposals error", err)
		return
	}
	for _, p := range pending {
		go c.queryStaking(ctx, p)
	}
}

func (s *stakingService) Init(ctx context.Context) *stakingService {
	rpc := g.Cfg().MustGet(ctx, "app.rpc").String()
	api, closer, err := client.NewFullNodeRPCV2(ctx, rpc, nil)
	if err != nil {
		g.Log().Error(ctx, "create full node rpc error", err)
		s.Err = err
		return s
	}
	defer closer()
	s.Api = api
	return s
}

func (s *stakingService) StartByPlan(ctx context.Context, p entity.SliverPlan) {
	c := s.Init(ctx)
	if c.Err != nil {
		g.Log().Error(ctx, "init staking service error", c.Err)
		return
	}
	c.queryStaking(ctx, p)
}
func (s *stakingService) queryStaking(ctx context.Context, p entity.SliverPlan) {
	cfStakingContract := g.Cfg().MustGet(ctx, "app.stakingContract").String()
	stakingContract, err := utility.ToEthAddress(cfStakingContract)
	if err != nil {
		g.Log().Error(ctx, "invalid staking contract address", err)
		return
	}
	stakingAbi, err := os.ReadFile("resource/staking-abi.json")
	if err != nil {
		g.Log().Error(ctx, "read staking-abi.json error", err)
		return
	}
	abi, err := abi.JSON(strings.NewReader(string(stakingAbi)))
	if err != nil {
		g.Log().Error(ctx, "parse staking-abi.json error", err)
		return
	}
	userAddr, err := utility.ToEthAddress(p.StakingAddress)
	if err != nil {
		g.Log().Error(ctx, "invalid user address", err)
		return
	}
	data, err := abi.Pack("getAllStakeInfo", userAddr)
	if err != nil {
		g.Log().Error(ctx, "pack getAllStakeInfo data error", err)
		return
	}
	// call contract
	latest := "latest"
	result, err := s.Api.EthCall(ctx, ethtypes.EthCall{
		To:   &stakingContract,
		Data: data,
	}, ethtypes.EthBlockNumberOrHash{
		PredefinedBlock: &latest,
	})
	if err != nil {
		g.Log().Error(ctx, "call getAllStakeInfo error", err)
		return
	}
	var stakeInfos []struct {
		StakeId   *big.Int
		Amount    *big.Int
		StartTime *big.Int
		LockTime  *big.Int
		Active    bool
	}
	err = abi.UnpackIntoInterface(&stakeInfos, "getAllStakeInfo", result)
	if err != nil {
		g.Log().Error(ctx, "unpack getAllStakeInfo result error", err)
		return
	}
	stakeTime := gtime.New(0)
	stakeSuccess := false
	stakeId := int64(0)
	for _, info := range stakeInfos {
		if !info.Active {
			continue
		}
		startTime := gtime.New(info.StartTime.Int64())
		if startTime.Before(p.CreatedAt) {
			continue
		}
		isHave, err := dao.SliverPlan.Ctx(ctx).Where(g.Map{
			"staking_id":      info.StakeId.Int64(),
			"staking_address": p.StakingAddress,
		}).Count()
		if err != nil {
			g.Log().Error(ctx, "check staking id error", err)
			continue
		}
		if isHave > 0 {
			continue
		}
		amountDecimal := decimal.NewFromBigInt(info.Amount, 0)
		fmt.Println(amountDecimal.String())
		g.Log().Info(ctx, "stake info:", info, startTime)
		stakeTime = startTime
		stakeSuccess = true
		stakeId = info.StakeId.Int64()
		break
	}
	if !stakeSuccess {
		// sleep 5s and retry
		time.Sleep(time.Second * 5)
		s.queryStaking(ctx, p)
		return
	}
	_, err = dao.SliverPlan.Ctx(ctx).Data(&g.Map{
		"staking_id":   stakeId,
		"staking_time": stakeTime,
		"status":       consts.ConstPlanStatusWaitingAllocation,
	}).Where("ID", p.Id).Update()
	if err != nil {
		g.Log().Error(ctx, "update proposal status error", err)
		time.Sleep(time.Second * 5)
		s.queryStaking(ctx, p)
		return
	}
	g.Log().Info(ctx, "stake success info:", stakeTime, stakeId)

	Allocator.Init(ctx).Send(ctx, model.PlanTask{
		Id:            p.Id,
		ClientAddress: p.ClientAddress,
		DataCap:       decimal.NewFromFloat(p.DataCap),
	})
}
