// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SliverPlanDao is the data access object for the table sliver_plan.
type SliverPlanDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns SliverPlanColumns // columns contains all the column names of Table for convenient usage.
}

// SliverPlanColumns defines and stores column names for the table sliver_plan.
type SliverPlanColumns struct {
	Id             string //
	PId            string //
	ClientAddress  string //
	DataCap        string //
	StakingAmount  string //
	StakingId      string //
	StakingDays    string //
	StakingAddress string //
	Status         string //
	AllocateTime   string //
	AllocateTx     string //
	CreatedAt      string //
	StakingTime    string //
}

// sliverPlanColumns holds the columns for the table sliver_plan.
var sliverPlanColumns = SliverPlanColumns{
	Id:             "ID",
	PId:            "p_id",
	ClientAddress:  "client_address",
	DataCap:        "data_cap",
	StakingAmount:  "staking_amount",
	StakingId:      "staking_id",
	StakingDays:    "staking_days",
	StakingAddress: "staking_address",
	Status:         "status",
	AllocateTime:   "allocate_time",
	AllocateTx:     "allocate_tx",
	CreatedAt:      "created_at",
	StakingTime:    "staking_time",
}

// NewSliverPlanDao creates and returns a new DAO object for table data access.
func NewSliverPlanDao() *SliverPlanDao {
	return &SliverPlanDao{
		group:   "default",
		table:   "sliver_plan",
		columns: sliverPlanColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SliverPlanDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SliverPlanDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SliverPlanDao) Columns() SliverPlanColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SliverPlanDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SliverPlanDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SliverPlanDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
