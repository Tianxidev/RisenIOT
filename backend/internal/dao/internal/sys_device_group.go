// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDeviceGroupDao is the data access object for table sys_device_group.
type SysDeviceGroupDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SysDeviceGroupColumns // columns contains all the column names of Table for convenient usage.
}

// SysDeviceGroupColumns defines and stores column names for table sys_device_group.
type SysDeviceGroupColumns struct {
	Id      string // 分组ID
	Name    string // 分组名称
	Remarks string // 备注
}

// sysDeviceGroupColumns holds the columns for table sys_device_group.
var sysDeviceGroupColumns = SysDeviceGroupColumns{
	Id:      "id",
	Name:    "name",
	Remarks: "remarks",
}

// NewSysDeviceGroupDao creates and returns a new DAO object for table data access.
func NewSysDeviceGroupDao() *SysDeviceGroupDao {
	return &SysDeviceGroupDao{
		group:   "default",
		table:   "sys_device_group",
		columns: sysDeviceGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysDeviceGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysDeviceGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysDeviceGroupDao) Columns() SysDeviceGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysDeviceGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysDeviceGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysDeviceGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
