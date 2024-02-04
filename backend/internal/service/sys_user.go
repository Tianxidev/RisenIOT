// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/api/v1/system"
	"backend/internal/model"
	"backend/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISysUser interface {
		GetCasBinUserPrefix() string
		NotCheckAuthAdminIds(ctx context.Context) *gset.Set
		GetAdminUserByUsernamePassword(ctx context.Context, req *system.UserLoginReq) (user *model.LoginUserRes, err error)
		// GetUserByUsername 通过用户名获取用户信息
		GetUserByUsername(ctx context.Context, userName string) (user *model.LoginUserRes, err error)
		// GetUserById 通过用户名获取用户信息
		GetUserById(ctx context.Context, id uint64) (user *model.LoginUserRes, err error)
		// LoginLog 记录登录日志
		LoginLog(ctx context.Context, params *model.LoginLogParams)
		UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error)
		// GetAdminRules 获取用户菜单数据
		GetAdminRules(ctx context.Context, userId uint64) (menuList []*model.UserMenus, permissions []string, err error)
		// GetAdminRole 获取用户角色
		GetAdminRole(ctx context.Context, userId uint64, allRoleList []*entity.SysRole) (roles []*entity.SysRole, err error)
		// GetAdminRoleIds 获取用户角色ids
		GetAdminRoleIds(ctx context.Context, userId uint64) (roleIds []uint, err error)
		GetAllMenus(ctx context.Context) (menus []*model.UserMenus, err error)
		GetAdminMenusByRoleIds(ctx context.Context, roleIds []uint) (menus []*model.UserMenus, err error)
		GetMenusTree(menus []*model.UserMenus, pid uint) []*model.UserMenus
		GetPermissions(ctx context.Context, roleIds []uint) (userButtons []string, err error)
		// List 用户列表
		List(ctx context.Context, req *system.UserSearchReq) (total interface{}, userList []*entity.SysUser, err error)
		// GetUsersRoleDept 获取多个用户角色 部门信息
		GetUsersRoleDept(ctx context.Context, userList []*entity.SysUser) (users []*model.SysUserRoleDeptRes, err error)
		Add(ctx context.Context, req *system.UserAddReq) (err error)
		Edit(ctx context.Context, req *system.UserEditReq) (err error)
		// AddUserPost 添加用户岗位信息
		AddUserPost(ctx context.Context, tx gdb.TX, postIds []int64, userId int64) (err error)
		// EditUserRole 修改用户角色信息
		EditUserRole(ctx context.Context, roleIds []int64, userId int64) (err error)
		UserNameOrMobileExists(ctx context.Context, userName, mobile string, id ...int64) error
		// GetEditUser 获取编辑用户信息
		GetEditUser(ctx context.Context, id uint64) (res *system.UserGetEditRes, err error)
		// GetUserInfoById 通过Id获取用户信息
		GetUserInfoById(ctx context.Context, id uint64, withPwd ...bool) (user *entity.SysUser, err error)
		// GetUserPostIds 获取用户岗位
		GetUserPostIds(ctx context.Context, userId uint64) (postIds []int64, err error)
		// ResetUserPwd 重置用户密码
		ResetUserPwd(ctx context.Context, req *system.UserResetPwdReq) (err error)
		ChangeUserStatus(ctx context.Context, req *system.UserStatusReq) (err error)
		// Delete 删除用户
		Delete(ctx context.Context, ids []int) (err error)
		// GetUsers 通过用户ids查询多个用户信息
		GetUsers(ctx context.Context, ids []int) (users []*model.SysUserSimpleRes, err error)
	}
)

var (
	localSysUser ISysUser
)

func SysUser() ISysUser {
	if localSysUser == nil {
		panic("implement not found for interface ISysUser, forgot register?")
	}
	return localSysUser
}

func RegisterSysUser(i ISysUser) {
	localSysUser = i
}
