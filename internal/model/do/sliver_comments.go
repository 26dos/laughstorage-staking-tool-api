// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SliverComments is the golang structure of table sliver_comments for DAO operations like Where/Data.
type SliverComments struct {
	g.Meta         `orm:"table:sliver_comments, do:true"`
	Id             interface{} //
	WalletAddress  interface{} //
	CommentContent interface{} //
	ParentId       interface{} //
	PId            interface{} //
	CreatedAt      *gtime.Time //
}
