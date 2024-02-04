// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	Id            uint64      `json:"id"            ` //
	UserName      string      `json:"userName"      ` // 用户名
	Mobile        string      `json:"mobile"        ` // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	UserNickname  string      `json:"userNickname"  ` // 用户昵称
	Birthday      int         `json:"birthday"      ` // 生日
	UserPassword  string      `json:"userPassword"  ` // 登录密码;cmf_password加密
	UserSalt      string      `json:"userSalt"      ` // 加密盐
	UserStatus    uint        `json:"userStatus"    ` // 用户状态;0:禁用,1:正常,2:未验证
	UserEmail     string      `json:"userEmail"     ` // 用户登录邮箱
	Sex           int         `json:"sex"           ` // 性别;0:保密,1:男,2:女
	Avatar        string      `json:"avatar"        ` // 用户头像
	DeptId        uint64      `json:"deptId"        ` // 部门id
	Remark        string      `json:"remark"        ` // 备注
	IsAdmin       int         `json:"isAdmin"       ` // 是否后台管理员 1 是  0   否
	Address       string      `json:"address"       ` // 联系地址
	Describe      string      `json:"describe"      ` // 描述信息
	LastLoginIp   string      `json:"lastLoginIp"   ` // 最后登录ip
	LastLoginTime *gtime.Time `json:"lastLoginTime" ` // 最后登录时间
	CreatedAt     *gtime.Time `json:"createdAt"     ` // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     ` // 更新时间
	DeletedAt     *gtime.Time `json:"deletedAt"     ` // 删除时间
}
