package token

import (
	"backend/internal/consts"
	"backend/internal/model"
	"backend/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tiger1103/gfast-token/gftoken"
)

type sToken struct {
	*gftoken.GfToken
}

func New() *sToken {
	var (
		ctx = gctx.New()
		opt *model.TokenOptions
		fun gftoken.OptionFunc
		err = g.Cfg().MustGet(ctx, "gfToken").Struct(&opt)
	)

	if err != nil {
		g.Log().Error(ctx, "gfToken配置文件读取失败", err)
	}

	if opt.CacheModel == consts.CacheModelRedis {
		fun = gftoken.WithGRedis()
	} else {
		fun = gftoken.WithGCache()
	}
	return &sToken{
		GfToken: gftoken.NewGfToken(
			gftoken.WithCacheKey(opt.CacheKey),
			gftoken.WithTimeout(opt.Timeout),
			gftoken.WithMaxRefresh(opt.MaxRefresh),
			gftoken.WithMultiLogin(opt.MultiLogin),
			gftoken.WithExcludePaths(opt.ExcludePaths),
			fun,
		),
	}
}

func init() {
	service.RegisterToken(New())
}

// Get 获取 GFToken 对象
func (s *sToken) Get() *gftoken.GfToken {
	return s.GfToken
}
