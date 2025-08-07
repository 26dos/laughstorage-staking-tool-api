// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SliverPlan is the golang structure of table sliver_plan for DAO operations like Where/Data.
type SliverPlan struct {
	g.Meta         `orm:"table:sliver_plan, do:true"`
	Id             interface{} //
	PId            interface{} //
	ClientAddress  interface{} //
	DataCap        interface{} //
	StakingAmount  interface{} //
	StakingId      interface{} //
	StakingDays    interface{} //
	StakingAddress interface{} //
	Status         interface{} //
	AllocateTime   *gtime.Time //
	AllocateTx     interface{} //
	CreatedAt      *gtime.Time //
	StakingTime    *gtime.Time //
}
