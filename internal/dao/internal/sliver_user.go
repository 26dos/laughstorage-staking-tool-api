// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SliverUserDao is the data access object for the table sliver_user.
type SliverUserDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns SliverUserColumns // columns contains all the column names of Table for convenient usage.
}

// SliverUserColumns defines and stores column names for the table sliver_user.
type SliverUserColumns struct {
	Id          string // 唯一自增ID
	LoginName   string // login name
	LoginPass   string //
	Role        string // role
	Wallet      string // 钱包地址
	DisplayName string // 显示名称
	Email       string // 邮箱地址
	KycStatus   string // kyc状态
	CreatedAt   string // 创建时间
	LastLoginAt string // 最后登录时间
}

// sliverUserColumns holds the columns for the table sliver_user.
var sliverUserColumns = SliverUserColumns{
	Id:          "ID",
	LoginName:   "login_name",
	LoginPass:   "login_pass",
	Role:        "role",
	Wallet:      "wallet",
	DisplayName: "display_name",
	Email:       "email",
	KycStatus:   "kyc_status",
	CreatedAt:   "created_at",
	LastLoginAt: "last_login_at",
}

// NewSliverUserDao creates and returns a new DAO object for table data access.
func NewSliverUserDao() *SliverUserDao {
	return &SliverUserDao{
		group:   "default",
		table:   "sliver_user",
		columns: sliverUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SliverUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SliverUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SliverUserDao) Columns() SliverUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SliverUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SliverUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SliverUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
