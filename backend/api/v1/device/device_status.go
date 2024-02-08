package device

import (
	commonApi "backend/api/v1/common"
	"backend/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// StatusSearchReq 分页请求参数
type StatusSearchReq struct {
	g.Meta             `path:"/deviceStatus/list" tags:"设备状态" method:"get" summary:"设备状态列表"`
	DeviceId           string `p:"deviceId"`           //设备ID
	Status             string `p:"status"`             //状态
	TimeOut            string `p:"timeOut"`            //超时时间
	UpTime             string `p:"upTime"`             //上线时间
	DownTime           string `p:"downTime"`           //离线时间
	LastDataUpdateTime string `p:"lastDataUpdateTime"` //最新一次数据更新时间
	BeginTime          string `p:"beginTime"`          //开始时间
	EndTime            string `p:"endTime"`            //结束时间
	commonApi.PageReq
}
type StatusSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.SysDeviceStatus `json:"list"`
}
type StatusGetReq struct {
	g.Meta `path:"/deviceStatus/get" tags:"设备状态" method:"get" summary:"获取设备状态数据"`
	Id     int `p:"id"`
}
type StatusGetRes struct {
	g.Meta `mime:"application/json"`
	*entity.SysDeviceStatus
}

// StatusAddReq 添加操作请求参数
type StatusAddReq struct {
	g.Meta             `path:"/deviceStatus/add" tags:"设备状态" method:"post" summary:"添加设备状态"`
	DeviceId           int  `p:"deviceId" v:"required#设备ID不能为空"`
	Status             int  `p:"status" v:"required#状态不能为空"`
	TimeOut            int  `p:"timeOut" v:"required#超时时间不能为空"`
	UpTime             int  `p:"upTime" `
	DownTime           int  `p:"downTime" `
	LastDataUpdateTime uint `p:"lastDataUpdateTime" v:"required#最新一次数据更新时间不能为空"`
}
type StatusAddRes struct {
}

// StatusEditReq 修改操作请求参数
type StatusEditReq struct {
	g.Meta             `path:"/deviceStatus/edit" tags:"设备状态" method:"put" summary:"修改设备状态"`
	Id                 int  `p:"id" v:"required#主键ID不能为空"`
	DeviceId           int  `p:"deviceId" v:"required#设备ID不能为空"`
	Status             int  `p:"status" v:"required#状态不能为空"`
	TimeOut            int  `p:"timeOut" v:"required#超时时间不能为空"`
	UpTime             int  `p:"upTime" `
	DownTime           int  `p:"downTime" `
	LastDataUpdateTime uint `p:"lastDataUpdateTime" v:"required#最新一次数据更新时间不能为空"`
}
type StatusEditRes struct {
}

// StatusStatusReq 设置用户状态参数
type StatusStatusReq struct {
	g.Meta `path:"/deviceStatus/status" tags:"设备状态" method:"put" summary:"修改设备状态状态"`
	Id     int `p:"id" v:"required#主键ID不能为空"`
	Status int `p:"status" v:"required#状态不能为空"`
}
type StatusStatusRes struct {
}
type StatusDeleteReq struct {
	g.Meta `path:"/deviceStatus/delete" tags:"设备状态" method:"delete" summary:"删除设备状态"`
	Ids    []int `p:"ids"`
}
type StatusDeleteRes struct {
}
