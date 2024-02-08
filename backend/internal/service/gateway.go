// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IGateway interface {
		// Start 网关启动
		Start(ctx context.Context)
	}
)

var (
	localGateway IGateway
)

func Gateway() IGateway {
	if localGateway == nil {
		panic("implement not found for interface IGateway, forgot register?")
	}
	return localGateway
}

func RegisterGateway(i IGateway) {
	localGateway = i
}
