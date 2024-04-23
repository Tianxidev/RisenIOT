/*
* @desc:在线用户
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/1/10 16:57
 */

package system

import (
	commonApi "backend/api/v1/common"
	"backend/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserOnlineSearchReq 列表搜索参数
type SysUserOnlineSearchReq struct {
	g.Meta   `path:"/system/online/list" tags:"在线用户管理" method:"get" summary:"列表"`
	Username string `p:"userName"`
	Ip       string `p:"ipaddr"`
	commonApi.PageReq
	commonApi.Author
}

// SysUserOnlineSearchRes 列表结果
type SysUserOnlineSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.SysUserOnline `json:"list"`
}

// SysUserOnlineForceLogoutReq 强制用户退出登录请求
type SysUserOnlineForceLogoutReq struct {
	g.Meta `path:"/system/online/forceLogout" tags:"在线用户管理" method:"delete" summary:"强制用户退出登录"`
	commonApi.Author
	Ids []int `p:"ids" v:"required#ids不能为空"`
}

// SysUserOnlineForceLogoutRes 强制用户退出登录响应
type SysUserOnlineForceLogoutRes struct {
	commonApi.EmptyRes
}

// SysUserOnlineForceLogoutAllReq 强制退出所有用户请求
type SysUserOnlineForceLogoutAllReq struct {
	g.Meta `path:"/system/online/forceLogoutAll" tags:"在线用户管理" method:"delete" summary:"强制退出所有用户"`
}

// SysUserOnlineForceLogoutAllRes 强制退出所有用户响应
type SysUserOnlineForceLogoutAllRes struct {
	commonApi.EmptyRes
}
