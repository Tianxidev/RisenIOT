// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RiseniotUser is the golang structure of table riseniot_user for DAO operations like Where/Data.
type RiseniotUser struct {
	g.Meta     `orm:"table:riseniot_user, do:true"`
	Uuid       interface{} // 用户uuid
	Username   interface{} // 用户名, 唯一索引
	Password   interface{} // 密码
	Nickname   interface{} // 昵称
	Role       interface{} // 用户角色
	CreateTime *gtime.Time // 创建时间
}
