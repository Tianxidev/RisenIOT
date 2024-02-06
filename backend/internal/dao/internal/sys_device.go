// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDeviceDao is the data access object for table sys_device.
type SysDeviceDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns SysDeviceColumns // columns contains all the column names of Table for convenient usage.
}

// SysDeviceColumns defines and stores column names for table sys_device.
type SysDeviceColumns struct {
	Id         string //
	DeviceName string // 设备名称
	DeviceUuid string // 设备UUID
	DeviceUser string // 设备所属用户
	DeviceType string // 设备类型
}

// sysDeviceColumns holds the columns for table sys_device.
var sysDeviceColumns = SysDeviceColumns{
	Id:         "id",
	DeviceName: "device_name",
	DeviceUuid: "device_uuid",
	DeviceUser: "device_user",
	DeviceType: "device_type",
}

// NewSysDeviceDao creates and returns a new DAO object for table data access.
func NewSysDeviceDao() *SysDeviceDao {
	return &SysDeviceDao{
		group:   "default",
		table:   "sys_device",
		columns: sysDeviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysDeviceDao) Columns() SysDeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysDeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
