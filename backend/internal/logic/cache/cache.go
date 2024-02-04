package cache

import (
	"backend/internal/consts"
	"backend/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tiger1103/gfast-cache/cache"
)

type sCache struct {
	*cache.GfCache
	prefix string
}

func init() {
	service.RegisterCache(New())
}

func New() *sCache {
	var (
		ctx            = gctx.New()
		cacheContainer *cache.GfCache
	)
	prefix := g.Cfg().MustGet(ctx, "system.cache.prefix").String()
	model := g.Cfg().MustGet(ctx, "system.cache.model").String()
	if model == consts.CacheModelRedis {
		// redis
		cacheContainer = cache.NewRedis(prefix)
	} else {
		// memory
		cacheContainer = cache.New(prefix)
	}
	return &sCache{
		GfCache: cacheContainer,
		prefix:  prefix,
	}
}

// GCache 获取缓存
func (s *sCache) GCache() *cache.GfCache {
	return s.GfCache
}
