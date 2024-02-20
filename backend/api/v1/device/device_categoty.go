package device

import (
	commonApi "backend/api/v1/common"
	"backend/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type CategorySearchReq struct {
	g.Meta `path:"/device/category/list" tags:"设备类别" method:"get" summary:"设备数据结构列表"`
	KindId int `p:"kindId"` // 数据模板的主键
	//Name      string `p:"name"`      // 数据名称
	//Mark      string `p:"mark"`      // 数据标识
	//DataType  string `p:"dataType"`  // 数据类型
	//Unit      string `p:"unit"`      // 数据单位
	//Ratio     string `p:"ratio"`     // 变比系数
	//Format    string `p:"format"`    // 格式化显示
	//HomeShow  string `p:"homeShow"`  // 首页是否展示
	//BeginTime string `p:"beginTime"` // 开始时间
	//EndTime   string `p:"endTime"`   // 结束时间
	commonApi.PageReq
}

type CategorySearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.SysDeviceCategory `json:"list"`
	Kind *entity.SysDeviceKind       `json:"kind"`
}

type CategoryGetReq struct {
	g.Meta `path:"/device/category/get" tags:"设备类别" method:"get" summary:"获取设备类别数据"`
	Id     int `p:"id"`
}

type CategoryGetRes struct {
	g.Meta `mime:"application/json"`
	*entity.SysDeviceCategory
}

type CategoryAddReq struct {
	g.Meta   `path:"/device/category/add" tags:"设备类别" method:"post" summary:"添加设备类别"`
	KindId   int    `p:"kindId" `
	Name     string `p:"name" v:"required#数据名称不能为空"`
	Mark     string `p:"mark" `
	DataType int    `p:"dataType" `
	Unit     string `p:"unit" `
	Ratio    string `p:"ratio" `
	Format   string `p:"format" `
	HomeShow int    `p:"homeShow" `
	Remark   string `p:"remark" `
}
type CategoryAddRes struct {
}

type CategoryEditReq struct {
	g.Meta   `path:"/device/category/edit" tags:"设备类别" method:"put" summary:"修改设备类别"`
	Id       int    `p:"id" v:"required#主键ID不能为空"`
	KindId   int    `p:"kindId" `
	Name     string `p:"name" v:"required#数据名称不能为空"`
	Mark     string `p:"mark" `
	DataType int    `p:"dataType" `
	Unit     string `p:"unit" `
	Ratio    string `p:"ratio" `
	Format   string `p:"format" `
	HomeShow int    `p:"homeShow" `
	Remark   string `p:"remark" `
}

type CategoryEditRes struct {
}

type CategoryDeleteReq struct {
	g.Meta `path:"/device/category/delete" tags:"设备类别" method:"delete" summary:"删除设备类别"`
	Ids    []int `p:"ids"`
}

type CategoryDeleteRes struct {
}
