package device

import (
	commonApi "backend/api/v1/common"
	"backend/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

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

type KindSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.SysDeviceKind `json:"list"`
}
