// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SliverProposals is the golang structure of table sliver_proposals for DAO operations like Where/Data.
type SliverProposals struct {
	g.Meta              `orm:"table:sliver_proposals, do:true"`
	Id                  interface{} // 唯一自增ID
	PId                 interface{} // 项目ID
	ClientAddress       interface{} //
	PName               interface{} // 项目名称
	PContent            interface{} // 项目内容
	PUser               interface{} // 用户ID
	Status              interface{} // 状态值
	ReasonRejection     interface{} //
	RequestDataCap      interface{} //
	DataCap             interface{} // 申领份额
	KycStatus           interface{} //
	KycVerificationTime *gtime.Time //
	CreatedAt           *gtime.Time // 创建时间
	UpdateAt            *gtime.Time // 修改时间
}
