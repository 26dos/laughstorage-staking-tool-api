// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SliverComments is the golang structure for table sliver_comments.
type SliverComments struct {
	Id             int64       `json:"id"              orm:"ID"              ` //
	WalletAddress  string      `json:"wallet_address"  orm:"wallet_address"  ` //
	CommentContent string      `json:"comment_content" orm:"comment_content" ` //
	ParentId       int64       `json:"parent_id"       orm:"parent_id"       ` //
	PId            int64       `json:"p_id"            orm:"p_id"            ` //
	CreatedAt      *gtime.Time `json:"created_at"      orm:"created_at"      ` //
}
