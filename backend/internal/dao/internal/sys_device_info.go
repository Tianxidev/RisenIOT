// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDeviceInfoDao is the data access object for table sys_device_info.
type SysDeviceInfoDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SysDeviceInfoColumns // columns contains all the column names of Table for convenient usage.
}

// SysDeviceInfoColumns defines and stores column names for table sys_device_info.
type SysDeviceInfoColumns struct {
	Id                 string //
	Name               string //
	Group              string //
	Sn                 string //
	Pwd                string //
	Kind               string //
	Logo               string //
	Monitor            string //
	Location           string //
	CreatedAt          string //
	UpdatedAt          string //
	Status             string //
	TimeOut            string //
	StatusId           string //
	UpTime             string // 上线时间
	DownTime           string // 离线时间
	LastDataUpdateTime string // 最新一次数据更新时间
	CreateBy           string // 创建人ID
}

// sysDeviceInfoColumns holds the columns for table sys_device_info.
var sysDeviceInfoColumns = SysDeviceInfoColumns{
	Id:                 "id",
	Name:               "name",
	Group:              "group",
	Sn:                 "sn",
	Pwd:                "pwd",
	Kind:               "kind",
	Logo:               "logo",
	Monitor:            "monitor",
	Location:           "location",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	Status:             "status",
	TimeOut:            "time_out",
	StatusId:           "status_id",
	UpTime:             "up_time",
	DownTime:           "down_time",
	LastDataUpdateTime: "last_data_update_time",
	CreateBy:           "create_by",
}

// NewSysDeviceInfoDao creates and returns a new DAO object for table data access.
func NewSysDeviceInfoDao() *SysDeviceInfoDao {
	return &SysDeviceInfoDao{
		group:   "default",
		table:   "sys_device_info",
		columns: sysDeviceInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysDeviceInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysDeviceInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysDeviceInfoDao) Columns() SysDeviceInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysDeviceInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysDeviceInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysDeviceInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
