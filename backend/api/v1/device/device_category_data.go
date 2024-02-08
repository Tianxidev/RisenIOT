package device

import (
	commonApi "backend/api/v1/common"
	"backend/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CategoryDataSearchReq 分页请求参数
type CategoryDataSearchReq struct {
	g.Meta     `path:"/device/category/data/list" tags:"设备数据" method:"get" summary:"设备数据列表"`
	CategoryId int     `p:"categoryId"` //数据类别主键
	DeviceId   int     `p:"deviceId"`   //设备主键
	DataInt    string  `p:"dataInt"`    //int型数据
	DataStr    string  `p:"dataStr"`    //字符串型数据
	DataDouble float64 `p:"dataDouble"` //double型数据
	BeginTime  string  `p:"beginTime"`  //开始时间
	EndTime    string  `p:"endTime"`    //结束时间
	Top        int     `p:"top"`        //最新的多少个
	commonApi.PageReq
}

type CategoryDataSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.SysDeviceCategoryData `json:"list"`
}

type CategoryDataRecentReq struct {
	g.Meta     `path:"/device/category/data/recent" tags:"设备数据" method:"get" summary:"设备最近数据列表"`
	CategoryId int    `p:"categoryId"` //数据类别主键
	DeviceId   int    `p:"deviceId"`   //设备主键
	BeginTime  string `p:"beginTime"`  //开始时间
	EndTime    string `p:"endTime"`    //结束时间
	commonApi.PageReq
}

type CategoryDataComm struct {
	Id   int         `orm:"id,primary" json:"id"`   // 主键
	Data interface{} `orm:"data" json:"data"`       // 字符串型数据
	Time *gtime.Time `orm:"created_at" json:"time"` // 修改时间
}

type CategoryDataRecentRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	Data []*CategoryDataComm
}

type CategoryDataHistoryReq struct {
	g.Meta     `path:"/device/category/data/history" tags:"设备数据" method:"get" summary:"设备历史数据列表"`
	CategoryId int    `p:"categoryId"` //数据类别主键
	DeviceId   int    `p:"deviceId"`   //设备主键
	BeginTime  string `p:"beginTime"`  //开始时间
	EndTime    string `p:"endTime"`    //结束时间
	commonApi.PageReq
}

type CategoryDataHistoryRes struct {
	g.Meta `mime:"application/json"`
	Total  int `json:"total"`
	Data   []*CategoryDataComm
}

type CategoryDataGetReq struct {
	g.Meta `path:"/device/category/data/get" tags:"设备数据" method:"get" summary:"获取设备数据"`
	Id     int `p:"id"`
}
type CategoryDataGetRes struct {
	g.Meta `mime:"application/json"`
	*entity.SysDeviceCategoryData
}

// CategoryDataAddReq 添加操作请求参数
type CategoryDataAddReq struct {
	g.Meta     `path:"/device/category/data/add" tags:"设备数据" method:"post" summary:"添加设备数据"`
	CategoryId int         `p:"categoryId" v:"required#数据类别主键不能为空"`
	DeviceId   int         `p:"deviceId" v:"required#设备主键不能为空"`
	DataInt    uint        `p:"dataInt" `
	DataStr    string      `p:"dataStr" `
	DataDouble float64     `p:"dataDouble" `
	CreatedAt  *gtime.Time `p:"time"` // 创建时间
}
type CategoryDataAddRes struct {
}

// CategoryDataEditReq 修改操作请求参数
type CategoryDataEditReq struct {
	g.Meta     `path:"/device/category/data/edit" tags:"设备数据" method:"put" summary:"修改设备数据"`
	Id         int     `p:"id" v:"required#主键ID不能为空"`
	CategoryId int     `p:"categoryId" v:"required#数据类别主键不能为空"`
	DeviceId   int     `p:"deviceId" v:"required#设备主键不能为空"`
	DataInt    uint    `p:"dataInt" `
	DataStr    string  `p:"dataStr" `
	DataDouble float64 `p:"dataDouble" `
}
type CategoryDataEditRes struct {
}
type CategoryDataDeleteReq struct {
	g.Meta `path:"/device/category/data/delete" tags:"设备数据" method:"delete" summary:"删除设备数据"`
	Ids    []int `p:"ids"`
}
type CategoryDataDeleteRes struct {
}
