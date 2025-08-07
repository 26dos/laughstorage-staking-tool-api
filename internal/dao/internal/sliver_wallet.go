// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SliverWalletDao is the data access object for the table sliver_wallet.
type SliverWalletDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns SliverWalletColumns // columns contains all the column names of Table for convenient usage.
}

// SliverWalletColumns defines and stores column names for the table sliver_wallet.
type SliverWalletColumns struct {
	Id          string // 唯一自增ID
	Wallet      string // 钱包地址
	DisplayName string // 显示名称
	Email       string // 邮箱地址
	KycStatus   string // kyc状态
	KycInfo     string // kyc信息
	CreatedAt   string // 创建时间
	LastLoginAt string // 最后登录时间
}

// sliverWalletColumns holds the columns for the table sliver_wallet.
var sliverWalletColumns = SliverWalletColumns{
	Id:          "ID",
	Wallet:      "wallet",
	DisplayName: "display_name",
	Email:       "email",
	KycStatus:   "kyc_status",
	KycInfo:     "kyc_info",
	CreatedAt:   "created_at",
	LastLoginAt: "last_login_at",
}

// NewSliverWalletDao creates and returns a new DAO object for table data access.
func NewSliverWalletDao() *SliverWalletDao {
	return &SliverWalletDao{
		group:   "default",
		table:   "sliver_wallet",
		columns: sliverWalletColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SliverWalletDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SliverWalletDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SliverWalletDao) Columns() SliverWalletColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SliverWalletDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SliverWalletDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SliverWalletDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
