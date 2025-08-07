package service

import (
	"api/internal/consts"
	"api/internal/dao"
	"api/internal/model"
	"api/utility"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

var Allocator = new(allocator)

type allocator struct {
	Client            *ethclient.Client
	Err               error
	AllocatorContract common.Address
	Abi               abi.ABI
	AllocatorAccount  common.Address
}

func (a *allocator) Init(ctx context.Context) *allocator {
	rpcURL := g.Cfg().MustGet(ctx, "app.rpc").String()
	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		g.Log().Error(ctx, "failed to connect rpc", err)
		a.Err = err
		return a
	}
	a.Client = client
	a.AllocatorContract = common.HexToAddress(g.Cfg().MustGet(ctx, "allocator.contract").String())
	a.AllocatorAccount = common.HexToAddress(g.Cfg().MustGet(ctx, "allocator.account").String())
	allocatorAbi, err := os.ReadFile("resource/allocator-abi.json")
	if err != nil {
		g.Log().Error(ctx, "read allocator-abi.json error", err)
		return a
	}
	a.Abi, err = abi.JSON(strings.NewReader(string(allocatorAbi)))
	if err != nil {
		g.Log().Error(ctx, "parse allocator-abi.json error", err)
		return a
	}
	return a
}

func (a *allocator) Start(ctx context.Context) {
	a.Init(ctx)
	if a.Err != nil {
		g.Log().Error(ctx, "init allocator service error", a.Err)
		return
	}
	ps := []model.PlanTask{}
	err := dao.SliverPlan.Ctx(ctx).Where("status", consts.ConstPlanStatusWaitingAllocation).Scan(&ps)
	if err != nil {
		g.Log().Error(ctx, "failed to get plan tasks", err)
		return
	}
	for _, p := range ps {
		// sleep 10 seconds
		time.Sleep(10 * time.Second)
		go a.Send(ctx, p)
	}
}

func (a *allocator) Send(ctx context.Context, p model.PlanTask) {
	client := a.Client
	abi := a.Abi
	// read config
	allocatorPrivateKey := g.Cfg().MustGet(ctx, "allocator.accountPrivateKey").String()

	// parse private key
	privKey, err := crypto.HexToECDSA(allocatorPrivateKey)
	if err != nil {
		g.Log().Error(ctx, "invalid private key", err)
		return
	}
	clientAddrBytes, err := utility.ToFilAddressBytes(ctx, p.ClientAddress)
	if err != nil {
		g.Log().Error(ctx, "resolve to id bytes error", err)
		return
	}
	// convert TiB amount to bytes
	// 1 TiB = 1024^4 bytes = 1,099,511,627,776 bytes
	tibToBytes, _ := new(big.Int).SetString("1099511627776", 10) // 1024^4
	amountTib := new(big.Int)
	amountTib.SetString(p.DataCap.String(), 10)
	amountBigInt := new(big.Int).Mul(amountTib, tibToBytes)

	// call allowance method to query the amount
	err = a.checkAllowance(ctx, amountBigInt)
	if err != nil {
		g.Log().Error(ctx, "check allowance error", err)
		return
	}
	// pack the method call data
	data, err := abi.Pack("addVerifiedClient", clientAddrBytes, amountBigInt)
	if err != nil {
		g.Log().Error(ctx, "pack addVerifiedClient data error", err)
		return
	}
	// get the sender address from private key
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		g.Log().Error(ctx, "error casting public key to ECDSA")
		return
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	err = a.checkAllocatorList(ctx)
	if err != nil {
		g.Log().Error(ctx, "permission check failed", err)
		return
	}
	g.Log().Info(ctx, "fromAddress", fromAddress.Hex())
	// get nonce
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		g.Log().Error(ctx, "get nonce error", err)
		return
	}
	// get gas price
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		g.Log().Error(ctx, "get gas price error", err)
		return
	}
	// estimate gas limit
	msg := ethereum.CallMsg{
		From: fromAddress,
		To:   &a.AllocatorContract,
		Data: data,
	}
	gasLimit, err := client.EstimateGas(ctx, msg)
	if err != nil {
		g.Log().Error(ctx, "estimate gas limit error", err.Error())
		return
	}
	// create transaction
	tx := types.NewTransaction(
		nonce,
		a.AllocatorContract,
		big.NewInt(0), // value is 0 for contract calls
		gasLimit,
		gasPrice,
		data,
	)
	// sign transaction
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		g.Log().Error(ctx, "get chain id error", err)
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		g.Log().Error(ctx, "sign transaction error", err)
		return
	}
	// send transaction
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		g.Log().Error(ctx, "send transaction error", err)
		return
	}
	txHash := signedTx.Hash().Hex()
	_, err = dao.SliverPlan.Ctx(ctx).Data(&g.Map{
		"allocate_time": gtime.Now().FormatTo("Y-m-d H:i:s"),
		"status":        consts.ConstPlanStatusSuccess,
		"allocate_tx":   txHash,
	}).Where("ID", p.Id).Update()
	if err != nil {
		g.Log().Error(ctx, "update plan status error", err)
		return
	}
	g.Log().Info(ctx, "addVerifiedClient transaction sent and success", "txHash", txHash)
}

func (a *allocator) checkAllowance(ctx context.Context, amountBigInt *big.Int) error {
	abi := a.Abi
	allowanceData, err := abi.Pack("allowance", a.AllocatorAccount)
	if err != nil {
		return err
	}
	msg := ethereum.CallMsg{
		To:   &a.AllocatorContract,
		Data: allowanceData,
	}
	result, err := a.Client.CallContract(ctx, msg, nil)
	if err != nil {
		return err
	}
	var allowanceAmount *big.Int
	err = abi.UnpackIntoInterface(&allowanceAmount, "allowance", result)
	if err != nil {
		return err
	}
	if allowanceAmount.Cmp(amountBigInt) < 0 {
		g.Log().Error(ctx, "insufficient allowance", "required", amountBigInt.String(), "available", allowanceAmount.String())
		return fmt.Errorf("insufficient allowance")
	}
	g.Log().Info(ctx, "current allowance", "allocator account:", a.AllocatorAccount.Hex(), "allowance:", allowanceAmount.String(), "required:", amountBigInt.String())
	return nil
}

func (a *allocator) checkAllocatorList(ctx context.Context) error {
	abi := a.Abi
	getAllocatorsData, err := abi.Pack("getAllocators")
	if err != nil {
		return err
	}

	msg := ethereum.CallMsg{
		To:   &a.AllocatorContract,
		Data: getAllocatorsData,
	}

	result, err := a.Client.CallContract(ctx, msg, nil)
	if err != nil {
		return err
	}

	var allocators []common.Address
	err = abi.UnpackIntoInterface(&allocators, "getAllocators", result)
	if err != nil {
		return err
	}

	isAllocator := slices.Contains(allocators, a.AllocatorAccount)
	if !isAllocator {
		g.Log().Error(ctx, "caller is not in allocators list", "caller", a.AllocatorAccount.Hex())
		return fmt.Errorf("caller is not in allocators list")
	}

	return nil
}
