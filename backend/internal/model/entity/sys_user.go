// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	Id            uint64      `json:"id"            orm:"id"              ` //
	UserName      string      `json:"userName"      orm:"user_name"       ` // 用户名
	Mobile        string      `json:"mobile"        orm:"mobile"          ` // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	UserNickname  string      `json:"userNickname"  orm:"user_nickname"   ` // 用户昵称
	Birthday      int         `json:"birthday"      orm:"birthday"        ` // 生日
	UserPassword  string      `json:"userPassword"  orm:"user_password"   ` // 登录密码;cmf_password加密
	UserSalt      string      `json:"userSalt"      orm:"user_salt"       ` // 加密盐
	UserStatus    uint        `json:"userStatus"    orm:"user_status"     ` // 用户状态;0:禁用,1:正常,2:未验证
	UserEmail     string      `json:"userEmail"     orm:"user_email"      ` // 用户登录邮箱
	Sex           int         `json:"sex"           orm:"sex"             ` // 性别;0:保密,1:男,2:女
	Avatar        string      `json:"avatar"        orm:"avatar"          ` // 用户头像
	DeptId        uint64      `json:"deptId"        orm:"dept_id"         ` // 部门id
	Remark        string      `json:"remark"        orm:"remark"          ` // 备注
	IsAdmin       int         `json:"isAdmin"       orm:"is_admin"        ` // 是否后台管理员 1 是  0   否
	Address       string      `json:"address"       orm:"address"         ` // 联系地址
	Describe      string      `json:"describe"      orm:"describe"        ` // 描述信息
	Pushdeer      string      `json:"pushdeer"      orm:"pushdeer"        ` // PushDeer推送Key
	LastLoginIp   string      `json:"lastLoginIp"   orm:"last_login_ip"   ` // 最后登录ip
	LastLoginTime *gtime.Time `json:"lastLoginTime" orm:"last_login_time" ` // 最后登录时间
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      ` // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      ` // 更新时间
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"      ` // 删除时间
}
