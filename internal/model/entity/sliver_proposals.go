// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SliverProposals is the golang structure for table sliver_proposals.
type SliverProposals struct {
	Id                  int64       `json:"id"                    orm:"ID"                    ` // 唯一自增ID
	PId                 string      `json:"p_id"                  orm:"p_id"                  ` // 项目ID
	ClientAddress       string      `json:"client_address"        orm:"client_address"        ` //
	PName               string      `json:"p_name"                orm:"p_name"                ` // 项目名称
	PContent            string      `json:"p_content"             orm:"p_content"             ` // 项目内容
	PUser               int64       `json:"p_user"                orm:"p_user"                ` // 用户ID
	Status              string      `json:"status"                orm:"status"                ` // 状态值
	ReasonRejection     string      `json:"reason_rejection"      orm:"reason_rejection"      ` //
	RequestDataCap      string      `json:"request_data_cap"      orm:"request_data_cap"      ` //
	DataCap             string      `json:"data_cap"              orm:"data_cap"              ` // 申领份额
	KycStatus           string      `json:"kyc_status"            orm:"kyc_status"            ` //
	KycVerificationTime *gtime.Time `json:"kyc_verification_time" orm:"kyc_verification_time" ` //
	CreatedAt           *gtime.Time `json:"created_at"            orm:"created_at"            ` // 创建时间
	UpdateAt            *gtime.Time `json:"update_at"             orm:"update_at"             ` // 修改时间
}
