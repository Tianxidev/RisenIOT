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
)

type (
	ISysUserOnline interface {
		Invoke(ctx context.Context, params *model.SysUserOnlineParams)
		// SaveOnline 保存用户在线状态
		SaveOnline(ctx context.Context, params *model.SysUserOnlineParams)
		// CheckUserOnline 检查在线用户
		CheckUserOnline(ctx context.Context)
		// GetOnlineListPage 搜素在线用户列表
		GetOnlineListPage(ctx context.Context, req *system.SysUserOnlineSearchReq, hasToken ...bool) (res *system.SysUserOnlineSearchRes, err error)
		UserIsOnline(ctx context.Context, token string) bool
		DeleteOnlineByToken(ctx context.Context, token string) (err error)
		ForceLogout(ctx context.Context, ids []int) (err error)
		GetInfosByIds(ctx context.Context, ids []int) (onlineList []*entity.SysUserOnline, err error)
	}
)

var (
	localSysUserOnline ISysUserOnline
)

func SysUserOnline() ISysUserOnline {
	if localSysUserOnline == nil {
		panic("implement not found for interface ISysUserOnline, forgot register?")
	}
	return localSysUserOnline
}

func RegisterSysUserOnline(i ISysUserOnline) {
	localSysUserOnline = i
}
