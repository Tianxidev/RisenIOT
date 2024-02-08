// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/api/v1/system"
	"backend/internal/model"
	"context"
)

type (
	ISysLoginLog interface {
		Invoke(ctx context.Context, data *model.LoginLogParams)
		List(ctx context.Context, req *system.LoginLogSearchReq) (res *system.LoginLogSearchRes, err error)
		DeleteLoginLogByIds(ctx context.Context, ids []int) (err error)
		ClearLoginLog(ctx context.Context) (err error)
	}
)

var (
	localSysLoginLog ISysLoginLog
)

func SysLoginLog() ISysLoginLog {
	if localSysLoginLog == nil {
		panic("implement not found for interface ISysLoginLog, forgot register?")
	}
	return localSysLoginLog
}

func RegisterSysLoginLog(i ISysLoginLog) {
	localSysLoginLog = i
}
