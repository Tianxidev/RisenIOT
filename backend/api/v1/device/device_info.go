package device

import (
	commonApi "backend/api/v1/common"
	"backend/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type InfoSearchReq struct {
	g.Meta    `path:"/device/info/list" tags:"设备信息" method:"get" summary:"设备信息列表"`
	Name      string      `p:"name"`      // 设备名称
	Group     string      `p:"group"`     // 设备组
	Sn        string      `p:"sn"`        // SN
	Pwd       string      `p:"pwd"`       // 密码
	Kind      string      `p:"kind"`      // 设备类别
	Logo      string      `p:"logo"`      // logo
	Monitor   string      `p:"monitor"`   // 是否监视
	Status    int         `p:"status"`    // 状态
	Location  string      `p:"location"`  // 地理位置
	CreatedAt *gtime.Time `p:"createdAt"` // 创建时间
	BeginTime string      `p:"beginTime"` // 开始时间
	EndTime   string      `p:"endTime"`   // 结束时间
	commonApi.PageReq
}

type InfoSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.SysDeviceInfo `json:"list"`
}

type InfoGetReq struct {
	g.Meta `path:"/device/info/get" tags:"设备信息" method:"get" summary:"获取设备信息数据"`
	Id     int `p:"id"`
}

type InfoGetRes struct {
	g.Meta       `mime:"application/json"`
	Info         *entity.SysDeviceInfo       `json:"info"`
	Kind         *entity.SysDeviceKind       `json:"kind"`
	CategoryList []*entity.SysDeviceCategory `json:"categoryList"`
}

type InfoAddReq struct {
	g.Meta   `path:"/device/info/add" tags:"设备信息" method:"post" summary:"添加设备信息"`
	Name     string `p:"name" v:"required#设备名称不能为空"`
	Group    int    `p:"group"`
	Sn       string `p:"sn"`
	Pwd      string `p:"pwd" `
	Kind     int    `p:"kind" v:"required#设备类别不能为空"`
	Logo     string `p:"logo" `
	Monitor  int    `p:"monitor" `
	Location int    `p:"location" `
}

type InfoAddRes struct {
}

type InfoEditReq struct {
	g.Meta   `path:"/device/info/edit" tags:"设备信息" method:"put" summary:"修改设备信息"`
	Id       int    `p:"id" v:"required#主键ID不能为空"`
	Name     string `p:"name" v:"required#设备名称不能为空"`
	Group    int    `p:"group" v:"required#设备组不能为空"`
	Sn       string `p:"sn" v:"required#SN不能为空"`
	Pwd      string `p:"pwd" `
	Kind     int    `p:"kind" v:"required#设备类别不能为空"`
	Logo     string `p:"logo" `
	Monitor  int    `p:"monitor" `
	Location int    `p:"location" `
}

type InfoEditRes struct {
}

type InfoDeleteReq struct {
	g.Meta `path:"/device/info/delete" tags:"设备信息" method:"delete" summary:"删除设备信息"`
	Ids    []int `p:"ids" v:"required#设备ID不能为空"`
}

type InfoDeleteRes struct {
}
