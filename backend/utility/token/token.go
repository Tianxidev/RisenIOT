package token

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var cfg = g.Cfg()

type loginFnType func(r *ghttp.Request) (string, interface{})

// getPaths 获取配置文件中的数组
func getPaths(ctx context.Context, path string) g.SliceStr {
	badAuth, _ := cfg.Get(ctx, path)
	dataSlice := badAuth.Array()
	authPaths := make(g.SliceStr, len(dataSlice))
	for i, path := range dataSlice {
		authPaths[i] = path.(string)
	}
	return authPaths
}

// Init 启动 gtoken
func Init(ctx context.Context, loginFn loginFnType) *gtoken.GfToken {

	// 缓存模式 1:redis 2:内存
	cacheMode, _ := cfg.Get(ctx, "tokenizer.cacheMode")

	// 多端登录
	multiLogin, _ := cfg.Get(ctx, "tokenizer.multiLogin")

	// 全局中间件
	globalMiddleware, _ := cfg.Get(ctx, "tokenizer.globalMiddleware")

	return &gtoken.GfToken{
		LoginPath:        "/system/login",
		LoginBeforeFunc:  loginFn,
		LogoutPath:       "/system/logout",
		AuthPaths:        getPaths(ctx, "tokenizer.authPaths"),
		AuthExcludePaths: getPaths(ctx, "tokenizer.authExcludePaths"),
		GlobalMiddleware: globalMiddleware.Bool(),
		CacheMode:        cacheMode.Int8(),
		MultiLogin:       multiLogin.Bool(),
	}

}
