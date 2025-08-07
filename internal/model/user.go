package model

import "github.com/gogf/gf/v2/os/gtime"

type UserInfo struct {
	LoginName   string      `json:"login_name"    orm:"login_name"    ` // login name
	Role        string      `json:"role"          orm:"role"          ` // role
	Wallet      string      `json:"wallet"        orm:"wallet"        ` // 钱包地址
	DisplayName string      `json:"display_name"  orm:"display_name"  ` // 显示名称
	Email       string      `json:"email"         orm:"email"         ` // 邮箱地址
	KycStatus   string      `json:"kyc_status"    orm:"kyc_status"    ` // kyc状态
	CreatedAt   *gtime.Time `json:"created_at"    orm:"created_at"    ` // 创建时间
	LastLoginAt *gtime.Time `json:"last_login_at" orm:"last_login_at" ` // 最后登录时间
}
