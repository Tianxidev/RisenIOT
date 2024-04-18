// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/api/v1/system"
	"context"
)

type (
	ISysPush interface {
		// TextMsgSendPushDeer 文本消息推送到 PushDeer
		TextMsgSendPushDeer(ctx context.Context, key string, text string) bool
		// SavePushDeerConfig 保存 PushDeer 配置
		SavePushDeerConfig(ctx context.Context, key string) bool
		// QueryPushDeerConfig 查询 PushDeer 配置
		QueryPushDeerConfig(ctx context.Context, res *system.PushQueryConfigToPushDeerRes)
	}
)

var (
	localSysPush ISysPush
)

func SysPush() ISysPush {
	if localSysPush == nil {
		panic("implement not found for interface ISysPush, forgot register?")
	}
	return localSysPush
}

func RegisterSysPush(i ISysPush) {
	localSysPush = i
}
