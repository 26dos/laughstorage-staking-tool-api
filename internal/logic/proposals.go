package logic

import (
	p_auth "api/api/proposals/auth"
	"api/api/proposals/manage"
	"api/api/proposals/user"
	"api/internal/consts"
	"api/internal/dao"
	"api/internal/model"
	"api/internal/model/entity"
	"api/internal/service"
	"api/utility"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

var Proposals = new(proposalsLogic)

type proposalsLogic struct {
}

func (l *proposalsLogic) Create(ctx context.Context, req *p_auth.ProposalsReq) (*g.Map, error) {
	if !utility.JustDataCapValue(req.RequestDataCap) {
		return nil, service.CustomError.ParameterError(ctx, "DataCap must be a number + unit (e.g 100TiB, 600.46GiB, 1.5PiB)")
	}
	if req.SubmitType != consts.ConstDraftStatus && req.SubmitType != consts.ConstSubmitStatus {
		return nil, service.CustomError.ParameterError(ctx, "Wrong form data!")
	}
	user, _ := service.User.GetCtxUser(ctx)
	dbProposal := entity.SliverProposals{}
	err := dao.SliverProposals.Ctx(ctx).Where("p_user", user.Id).Scan(&dbProposal)
	isCreate := false
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			isCreate = true
		} else {
			g.Log().Error(ctx, "get proposal error", err)
			return nil, consts.ServerErr
		}
	}
	// if proposal is not create and status is not draft or reject, return error
	if !isCreate && (dbProposal.Status != consts.ConstDraftStatus || dbProposal.Status != consts.ConstRejectStatus) {
		return nil, service.CustomError.ParameterError(ctx, "this proposal cannot be edited at this time!")
	}
	formOrigin, err := os.ReadFile("resource/form.json")
	if err != nil {
		g.Log().Error(ctx, "read form.json error", err)
		return nil, consts.ServerErr
	}
	formModel := []model.Form{}
	err = json.Unmarshal(formOrigin, &formModel)
	if err != nil {
		g.Log().Error(ctx, "unmarshal form.json error", err)
		return nil, consts.ServerErr
	}
	formData := []model.ReqForm{}
	err = json.Unmarshal([]byte(req.Data), &formData)
	if err != nil {
		g.Log().Error(ctx, "unmarshal req.Data error", err)
		return nil, consts.ServerErr
	}
	submitType := req.SubmitType
	clientAddress := ""
	for _, v := range formModel {
		for _, field := range v.Fields {
			var found bool
			var formItem model.ReqForm
			var formValue string
			for _, item := range formData {
				if item.Key == field.Key {
					formItem = item
					found = true
					break
				}
			}
			if !field.Required && !found {
				continue
			}
			if field.Required && submitType != consts.ConstDraftStatus {
				if !found {
					return nil, service.CustomError.ParameterError(ctx, "Please enter the "+field.Label)
				}
				formValue = fmt.Sprintf("%v", formItem.Value)
				if utility.EmptyString(formValue) {
					return nil, service.CustomError.ParameterError(ctx, "Please select the "+field.Label)
				}
			}
			if field.Type == "select" {
				if utility.EmptyString(formValue) && (submitType == consts.ConstDraftStatus || !field.Required) {
					continue
				}
				if field.Key == "data_owner_country_region" || field.Key == "data_preparer_country_region" {
					if !slices.Contains(consts.CountryNames, formValue) {
						return nil, service.CustomError.ParameterError(ctx, "Please select the "+field.Label+" from the options")
					}
					continue
				}
				options := make([]string, len(field.Options))
				for i, opt := range field.Options {
					options[i] = fmt.Sprintf("%v", opt)
				}
				if field.Multiple {
					formValues := strings.Split(formValue, ",")
					for _, _value := range formValues {
						if !slices.Contains(options, _value) {
							return nil, service.CustomError.ParameterError(ctx, "Please select the "+field.Label+" from the options")
						}
					}
				} else {
					if !slices.Contains(options, formValue) {
						return nil, service.CustomError.ParameterError(ctx, "Please select the "+field.Label+" from the options")
					}
				}
			}
			if field.Key == "allocation_address" {
				err := utility.CheckValidAndActivatedF1Address(ctx, formValue)
				if err != nil {
					return nil, service.CustomError.ParameterError(ctx, fmt.Sprintf("Error in %s field: %s", field.Label, err.Error()))
				}
				clientAddress = formValue
			}
		}
	}
	formJson, err := json.Marshal(formData)
	if err != nil {
		return nil, consts.ServerErr
	}
	dbData := g.Map{
		"p_name":           req.PName,
		"p_content":        formJson,
		"p_user":           user.Id,
		"status":           req.SubmitType,
		"request_data_cap": req.RequestDataCap,
		"reason_rejection": "",
		"client_address":   clientAddress,
	}
	if isCreate {
		pid := uuid.New().String()
		dbData["p_id"] = pid
		_, err = dao.SliverProposals.Ctx(ctx).Data(&dbData).Insert()
		if err != nil {
			g.Log().Error(ctx, "insert proposal error", err)
			return nil, consts.ServerErr
		}
		return &g.Map{"pid": pid, "kyc_status": user.KycStatus}, nil
	}
	_, err = dao.SliverProposals.Ctx(ctx).Data(&dbData).Where("p_id", dbProposal.PId).Update()
	if err != nil {
		g.Log().Error(ctx, "update proposal error", err)
		return nil, consts.ServerErr
	}
	dao.SliverUser.Ctx(ctx).Where("id", user.Id).Data(g.Map{"display_name": req.PName}).Update()
	return &g.Map{"pid": dbProposal.PId, "kyc_status": user.KycStatus}, nil
}

func (l *proposalsLogic) GetMyProposals(ctx context.Context) (*model.Proposal, error) {
	user, _ := service.User.GetCtxUser(ctx)
	proposals := model.Proposal{}
	err := dao.SliverProposals.Ctx(ctx).Where("p_user", user.Id).WithAll().Scan(&proposals)
	if err != nil {
		return nil, service.CustomError.NoData(ctx, "No data found")
	}
	return &proposals, nil
}

func (l *proposalsLogic) GetProposals(ctx context.Context, req *user.ProposalsReq) (*model.PagedData, error) {
	proposals := []model.Proposal{}
	err := dao.SliverProposals.Ctx(ctx).Order("created_at desc").WithAll().Page(req.Page, req.Limit).Scan(&proposals)
	if err != nil {
		return nil, service.CustomError.NoData(ctx, "No data found")
	}
	return &model.PagedData{
		Paged: model.Paged{
			Total:     len(proposals),
			Page:      req.Page,
			Size:      req.Limit,
			TotalPage: len(proposals) / req.Limit,
		},
		Data: proposals,
	}, nil
}

func (l *proposalsLogic) GetProposalDetail(ctx context.Context, req *user.ProposalDetailReq) (*model.Proposal, error) {
	proposal := model.Proposal{}
	err := dao.SliverProposals.Ctx(ctx).Where("p_id", req.PId).WithAll().Scan(&proposal)
	if err != nil {
		return nil, service.CustomError.NoData(ctx, "No data found")
	}
	return &proposal, nil
}

func (l *proposalsLogic) AuditProposal(ctx context.Context, req *manage.AuditProposalReq) (*model.Proposal, error) {
	proposal := model.Proposal{}
	err := dao.SliverProposals.Ctx(ctx).Where("p_id", req.PID).Scan(&proposal)
	if err != nil {
		return nil, consts.ServerErr
	}
	if proposal.Status != consts.ConstSubmitStatus {
		return nil, service.CustomError.ParameterError(ctx, "The current proposal status cannot be reviewed!")
	}
	if proposal.KycStatus != consts.ConstKycStatusYes {
		g.Log().Error(ctx, "The current proposal has not completed kyc verification!", proposal.KycStatus, proposal)
		return nil, service.CustomError.ParameterError(ctx, "The current proposal has not completed kyc verification!")
	}
	_, err = dao.SliverProposals.Ctx(ctx).Data(&g.Map{
		"status":   consts.ConstProposalSuccessStatus,
		"data_cap": req.DataCap,
	}).Where("p_id", req.PID).Update()
	if err != nil {
		return nil, consts.ServerErr
	}
	err = dao.SliverProposals.Ctx(ctx).Where("p_id", req.PID).WithAll().Scan(&proposal)
	if err != nil {
		return nil, consts.ServerErr
	}
	return &proposal, nil
}
func (l *proposalsLogic) RejectionProposal(ctx context.Context, req *manage.RejectionReq) (*model.Proposal, error) {
	proposal := model.Proposal{}
	err := dao.SliverProposals.Ctx(ctx).Where("p_id", req.PID).Scan(&proposal)
	if err != nil {
		return nil, consts.ServerErr
	}
	if proposal.Status != consts.ConstSubmitStatus || proposal.KycStatus != consts.ConstKycStatusYes {
		return nil, service.CustomError.ParameterError(ctx, "The current proposal is not under review!")
	}
	_, err = dao.SliverProposals.Ctx(ctx).Data(&g.Map{
		"status":           consts.ConstRejectStatus,
		"reason_rejection": req.Reason,
	}).Where("p_id", req.PID).Update()
	if err != nil {
		return nil, consts.ServerErr
	}
	err = dao.SliverProposals.Ctx(ctx).Where("p_id", req.PID).WithAll().Scan(&proposal)
	return &proposal, nil
}

func (l *proposalsLogic) CreatePlan(ctx context.Context, req *manage.CreatePlanReq) (*model.Proposal, error) {
	proposal := entity.SliverProposals{}
	err := dao.SliverProposals.Ctx(ctx).Where("p_id", req.PID).Scan(&proposal)
	if err != nil {
		return nil, consts.ServerErr
	}
	if proposal.Status != consts.ConstProposalSuccessStatus {
		return nil, service.CustomError.ParameterError(ctx, "The current proposal is not under review!")
	}
	user := entity.SliverUser{}
	err = dao.SliverUser.Ctx(ctx).Where("ID", proposal.PUser).Scan(&user)
	if err != nil {
		return nil, consts.ServerErr
	}
	dataCap, _ := req.DataCap.Float64()
	amount, _ := req.Amount.Float64()
	plan := entity.SliverPlan{
		PId:            proposal.Id,
		DataCap:        dataCap,
		StakingAmount:  amount,
		StakingDays:    req.StakeDays,
		Status:         consts.ConstPlanStatusPending,
		StakingAddress: user.Wallet,
		ClientAddress:  proposal.ClientAddress,
		StakingId:      -1,
	}
	planId, err := dao.SliverPlan.Ctx(ctx).Data(&plan).InsertAndGetId()
	if err != nil {
		return nil, consts.ServerErr
	}
	res := model.Proposal{}
	err = dao.SliverProposals.Ctx(ctx).Where("p_id", req.PID).WithAll().Scan(&res)
	if err != nil {
		return nil, consts.ServerErr
	}
	plan.Id = planId
	go service.Staking.StartByPlan(context.Background(), plan)
	return &res, nil
}
