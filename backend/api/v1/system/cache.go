package system

import (
	commonApi "backend/api/v1/common"
	"github.com/gogf/gf/v2/frame/g"
)

type CacheRemoveReq struct {
	g.Meta `path:"/system/cache/remove" tags:"缓存管理" method:"delete" summary:"清除缓存"`
	commonApi.Author
}

type CacheRemoveRes struct {
	commonApi.EmptyRes
}
