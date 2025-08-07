// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SliverProposalsDao is the data access object for the table sliver_proposals.
type SliverProposalsDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of the current DAO.
	columns SliverProposalsColumns // columns contains all the column names of Table for convenient usage.
}

// SliverProposalsColumns defines and stores column names for the table sliver_proposals.
type SliverProposalsColumns struct {
	Id                  string // 唯一自增ID
	PId                 string // 项目ID
	ClientAddress       string //
	PName               string // 项目名称
	PContent            string // 项目内容
	PUser               string // 用户ID
	Status              string // 状态值
	ReasonRejection     string //
	RequestDataCap      string //
	DataCap             string // 申领份额
	KycStatus           string //
	KycVerificationTime string //
	CreatedAt           string // 创建时间
	UpdateAt            string // 修改时间
}

// sliverProposalsColumns holds the columns for the table sliver_proposals.
var sliverProposalsColumns = SliverProposalsColumns{
	Id:                  "ID",
	PId:                 "p_id",
	ClientAddress:       "client_address",
	PName:               "p_name",
	PContent:            "p_content",
	PUser:               "p_user",
	Status:              "status",
	ReasonRejection:     "reason_rejection",
	RequestDataCap:      "request_data_cap",
	DataCap:             "data_cap",
	KycStatus:           "kyc_status",
	KycVerificationTime: "kyc_verification_time",
	CreatedAt:           "created_at",
	UpdateAt:            "update_at",
}

// NewSliverProposalsDao creates and returns a new DAO object for table data access.
func NewSliverProposalsDao() *SliverProposalsDao {
	return &SliverProposalsDao{
		group:   "default",
		table:   "sliver_proposals",
		columns: sliverProposalsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SliverProposalsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SliverProposalsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SliverProposalsDao) Columns() SliverProposalsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SliverProposalsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SliverProposalsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SliverProposalsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
