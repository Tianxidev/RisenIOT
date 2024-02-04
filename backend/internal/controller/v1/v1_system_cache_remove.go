package v1

import (
	"backend/internal/consts"
	"backend/internal/service"
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"backend/api/v1/system"
)

func (c *ControllerSystem) CacheRemove(ctx context.Context, req *system.CacheRemoveReq) (res *system.CacheRemoveRes, err error) {
	service.Cache().GCache().RemoveByTag(ctx, consts.CacheSysDictTag)
	service.Cache().GCache().RemoveByTag(ctx, consts.CacheSysConfigTag)
	service.Cache().GCache().RemoveByTag(ctx, consts.CacheSysAuthTag)
	cacheRedis := g.Cfg().MustGet(ctx, "system.cache.model").String()
	if cacheRedis == consts.CacheModelRedis {
		cursor := 0
		cachePrefix := g.Cfg().MustGet(ctx, "system.cache.prefix").String()
		for {
			var v *gvar.Var
			v, err = g.Redis().Do(ctx, "scan", cursor, "match", cachePrefix+"*", "count", "100")
			if err != nil {
				return
			}
			data := gconv.SliceAny(v)
			var dataSlice []string
			err = gconv.Structs(data[1], &dataSlice)
			if err != nil {
				return
			}
			for _, d := range dataSlice {
				_, err = g.Redis().Do(ctx, "del", d)
				if err != nil {
					return
				}
			}
			cursor = gconv.Int(data[0])
			if cursor == 0 {
				break
			}
		}
	}
	return
}
