// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast-token/gftoken"
)

type (
	IToken interface {
		// Get 获取 GFToken 对象
		Get() *gftoken.GfToken
		IsLogin(r *ghttp.Request) (b bool)
		Middleware(group *ghttp.RouterGroup) error
	}
)

var (
	localToken IToken
)

func Token() IToken {
	if localToken == nil {
		panic("implement not found for interface IToken, forgot register?")
	}
	return localToken
}

func RegisterToken(i IToken) {
	localToken = i
}
