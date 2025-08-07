// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SliverUser is the golang structure of table sliver_user for DAO operations like Where/Data.
type SliverUser struct {
	g.Meta      `orm:"table:sliver_user, do:true"`
	Id          interface{} // 唯一自增ID
	LoginName   interface{} // login name
	LoginPass   interface{} //
	Role        interface{} // role
	Wallet      interface{} // 钱包地址
	DisplayName interface{} // 显示名称
	Email       interface{} // 邮箱地址
	KycStatus   interface{} // kyc状态
	CreatedAt   *gtime.Time // 创建时间
	LastLoginAt *gtime.Time // 最后登录时间
}
