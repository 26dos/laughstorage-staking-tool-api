package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type Comment struct {
	g.Meta         `orm:"table:sliver_comments"`
	Id             int64          `json:"id"              orm:"ID"              ` //
	WalletAddress  string         `json:"wallet_address"  orm:"wallet_address"  ` //
	CommentContent string         `json:"comment_content" orm:"comment_content" ` //
	ParentId       int64          `json:"parent_id"       orm:"parent_id"       ` //
	PId            int64          `json:"p_id"            orm:"p_id"            ` //
	User           User           `json:"user"            orm:"with:wallet=wallet_address"`
	Parent         *ParentComment `json:"parent"          orm:"with:ID=parent_id"`
	CreatedAt      *gtime.Time    `json:"created_at"      orm:"created_at"      ` //
}

type ParentComment struct {
	g.Meta         `orm:"table:sliver_comments"`
	Id             int64       `json:"id"              orm:"ID"              ` //
	WalletAddress  string      `json:"wallet_address"  orm:"wallet_address"  ` //
	CommentContent string      `json:"comment_content" orm:"comment_content" ` //
	ParentId       int64       `json:"parent_id"       orm:"parent_id"       ` //
	PId            int64       `json:"p_id"            orm:"p_id"            ` //
	User           User        `json:"user"            orm:"with:wallet=wallet_address"`
	CreatedAt      *gtime.Time `json:"created_at"      orm:"created_at"      ` //
}

type User struct {
	g.Meta      `orm:"table:sliver_user"`
	ID          int    `json:"id" orm:"id"`
	Wallet      string `json:"wallet" orm:"wallet"`
	DisplayName string `json:"display_name" orm:"display_name"`
}
