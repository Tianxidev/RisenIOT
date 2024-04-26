// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ISysCron interface {
		// StartBaseTask 启动基础任务
		StartBaseTask(ctx context.Context)
	}
)

var (
	localSysCron ISysCron
)

func SysCron() ISysCron {
	if localSysCron == nil {
		panic("implement not found for interface ISysCron, forgot register?")
	}
	return localSysCron
}

func RegisterSysCron(i ISysCron) {
	localSysCron = i
}
