package device

import (
	commonApi "backend/api/v1/common"
	"backend/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// KindSearchReq 产品类型列表请求参数
type KindSearchReq struct {
	g.Meta    `path:"/device/kind/list" tags:"设备产品管理" method:"get" summary:"产品类型列表"`
	Id        int    `p:"id"`        // Id
	Name      string `p:"name"`      // 设备品牌名称
	Mark      string `p:"mark"`      // 设备品牌标记
	TimeOut   string `p:"timeOut"`   // 超时时间
	BeginTime string `p:"beginTime"` // 开始时间
	EndTime   string `p:"endTime"`   // 结束时间
	commonApi.PageReq
}

// KindSearchRes 产品类型列表响应参数
type KindSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.SysDeviceKind `json:"list"`
}

// KindAddReq 添加产品类型请求参数
type KindAddReq struct {
	g.Meta  `path:"/device/kind/add" tags:"设备产品管理" method:"post" summary:"添加产品类型"`
	Name    string `p:"name" v:"required#设备品牌名称不能为空"`
	Mark    string `p:"mark" v:"required#设备品牌标记不能为空"`
	TimeOut int    `p:"timeOut" v:"required#超时时间不能为空"`
}

// KindAddRes 添加产品类型响应参数
type KindAddRes struct {
}

// KindEditReq 编辑产品类型请求参数
type KindEditReq struct {
	g.Meta  `path:"/device/kind/edit" tags:"设备产品管理" method:"put" summary:"编辑产品类型"`
	Id      int    `p:"id" v:"required#主键ID不能为空"`
	Name    string `p:"name" v:"required#设备品牌名称不能为空"`
	Mark    string `p:"mark" v:"required#设备品牌标记不能为空"`
	TimeOut int    `p:"timeOut" v:"required#超时时间不能为空"`
}

// KindEditRes 编辑产品类型响应参数
type KindEditRes struct {
}

// KindDeleteReq 删除产品类型请求参数
type KindDeleteReq struct {
	g.Meta `path:"/device/kind/delete" tags:"设备产品管理" method:"delete" summary:"删除产品类型"`
	Ids    []int `p:"ids" v:"required#主键ID不能为空"`
}

// KindDeleteRes 删除产品类型响应参数
type KindDeleteRes struct {
}
