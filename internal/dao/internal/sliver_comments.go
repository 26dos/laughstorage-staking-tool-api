// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SliverCommentsDao is the data access object for the table sliver_comments.
type SliverCommentsDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns SliverCommentsColumns // columns contains all the column names of Table for convenient usage.
}

// SliverCommentsColumns defines and stores column names for the table sliver_comments.
type SliverCommentsColumns struct {
	Id             string //
	WalletAddress  string //
	CommentContent string //
	ParentId       string //
	PId            string //
	CreatedAt      string //
}

// sliverCommentsColumns holds the columns for the table sliver_comments.
var sliverCommentsColumns = SliverCommentsColumns{
	Id:             "ID",
	WalletAddress:  "wallet_address",
	CommentContent: "comment_content",
	ParentId:       "parent_id",
	PId:            "p_id",
	CreatedAt:      "created_at",
}

// NewSliverCommentsDao creates and returns a new DAO object for table data access.
func NewSliverCommentsDao() *SliverCommentsDao {
	return &SliverCommentsDao{
		group:   "default",
		table:   "sliver_comments",
		columns: sliverCommentsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SliverCommentsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SliverCommentsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SliverCommentsDao) Columns() SliverCommentsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SliverCommentsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SliverCommentsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SliverCommentsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
