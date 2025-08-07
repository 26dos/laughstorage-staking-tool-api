// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SliverPlan is the golang structure for table sliver_plan.
type SliverPlan struct {
	Id             int64       `json:"id"              orm:"ID"              ` //
	PId            int64       `json:"p_id"            orm:"p_id"            ` //
	ClientAddress  string      `json:"client_address"  orm:"client_address"  ` //
	DataCap        float64     `json:"data_cap"        orm:"data_cap"        ` //
	StakingAmount  float64     `json:"staking_amount"  orm:"staking_amount"  ` //
	StakingId      int64       `json:"staking_id"      orm:"staking_id"      ` //
	StakingDays    int         `json:"staking_days"    orm:"staking_days"    ` //
	StakingAddress string      `json:"staking_address" orm:"staking_address" ` //
	Status         string      `json:"status"          orm:"status"          ` //
	AllocateTime   *gtime.Time `json:"allocate_time"   orm:"allocate_time"   ` //
	AllocateTx     string      `json:"allocate_tx"     orm:"allocate_tx"     ` //
	CreatedAt      *gtime.Time `json:"created_at"      orm:"created_at"      ` //
	StakingTime    *gtime.Time `json:"staking_time"    orm:"staking_time"    ` //
}
