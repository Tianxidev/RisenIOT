package system

import (
	commonApi "backend/api/v1/common"
	"github.com/gogf/gf/v2/frame/g"
)

type UserLoginOutReq struct {
	g.Meta `path:"/system/logout" tags:"登录" method:"get" summary:"退出登录"`
	commonApi.Author
}

type UserLoginOutRes struct {
	Ok bool `json:"ok"`
}
