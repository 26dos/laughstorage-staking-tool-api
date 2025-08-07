package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
)

type Form struct {
	Group  string  `json:"group"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Label    string        `json:"label"`
	Type     string        `json:"type"`
	Required bool          `json:"required"`
	Options  []interface{} `json:"options"`
	Key      string        `json:"key"`
	Hint     string        `json:"hint,omitempty"`
	Data     string        `json:"data,omitempty"`
	Multiple bool          `json:"multiple,omitempty"`
}

type ReqForm struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type Proposal struct {
	g.Meta              `orm:"table:sliver_proposals"`
	Id                  int64       `json:"id"             orm:"ID"             `               //
	PId                 string      `json:"p_id"           orm:"p_id"           `               // 项目ID
	PName               string      `json:"p_name"         orm:"p_name"         `               // 项目名称
	PContent            string      `json:"p_content"      orm:"p_content"      `               // 项目内容
	Status              string      `json:"status"         orm:"status"         `               // 状态值
	ReasonRejection     string      `json:"reason_rejection"      orm:"reason_rejection"      ` //
	StakingId           int         `json:"staking_id"     orm:"staking_id"     `               // 质押ID
	PUser               int         `json:"p_user"         orm:"p_user"         `               // 用户ID
	RequestDataCap      string      `json:"request_data_cap" orm:"request_data_cap" `           //
	DataCap             string      `json:"data_cap"       orm:"data_cap"       `               // 申领份额
	KycStatus           string      `json:"kyc_status"            orm:"kyc_status"            ` //
	Wallet              Wallet      `json:"wallet" orm:"with:ID=p_user"`
	Plans               []Plan      `json:"plans" orm:"with:p_id=id"`
	KycVerificationTime *gtime.Time `json:"kyc_verification_time" orm:"kyc_verification_time" ` //
	CreatedAt           *gtime.Time `json:"created_at"     orm:"created_at"     `               // 创建时间
	UpdateAt            *gtime.Time `json:"update_at"      orm:"update_at"      `               // 修改时间

}

type Plan struct {
	g.Meta        `orm:"table:sliver_plan"`
	PId           int64           `json:"p_id"           orm:"p_id"           `   //
	ClientAddress string          `json:"client_address" orm:"client_address" `   //
	DataCap       decimal.Decimal `json:"data_cap"       orm:"data_cap"       `   //
	StakingAmount decimal.Decimal `json:"staking_amount" orm:"staking_amount" `   //
	StakingId     int64           `json:"staking_id"     orm:"staking_id"     `   //
	StakingDays   int             `json:"staking_days"   orm:"staking_days"   `   //
	Status        string          `json:"status"         orm:"status"         `   //
	AllocateTime  *gtime.Time     `json:"allocate_time"  orm:"allocate_time"  `   //
	AllocateTx    string          `json:"allocate_tx"     orm:"allocate_tx"     ` //
	CreatedAt     *gtime.Time     `json:"created_at"     orm:"created_at"     `   //
	StakingTime   *gtime.Time     `json:"staking_time"   orm:"staking_time"   `   //
}

type PlanTask struct {
	Id            int64           `json:"id"              orm:"ID"              ` //
	ClientAddress string          `json:"client_address" orm:"client_address" `   //
	DataCap       decimal.Decimal `json:"data_cap"       orm:"data_cap"       `   //
}

type Wallet struct {
	g.Meta      `orm:"table:sliver_user"`
	ID          int    `json:"id" orm:"id"`
	Wallet      string `json:"wallet" orm:"wallet"`
	DisplayName string `json:"display_name" orm:"display_name"`
}
