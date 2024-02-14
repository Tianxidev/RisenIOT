// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDeviceKindDao is the data access object for table sys_device_kind.
type SysDeviceKindDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SysDeviceKindColumns // columns contains all the column names of Table for convenient usage.
}

// SysDeviceKindColumns defines and stores column names for table sys_device_kind.
type SysDeviceKindColumns struct {
	Id        string // 产品类型ID
	Name      string // 产品名称
	Mark      string // 产品标记
	TimeOut   string //
	CreatedAt string //
	UpdatedAt string //
}

// sysDeviceKindColumns holds the columns for table sys_device_kind.
var sysDeviceKindColumns = SysDeviceKindColumns{
	Id:        "id",
	Name:      "name",
	Mark:      "mark",
	TimeOut:   "time_out",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysDeviceKindDao creates and returns a new DAO object for table data access.
func NewSysDeviceKindDao() *SysDeviceKindDao {
	return &SysDeviceKindDao{
		group:   "default",
		table:   "sys_device_kind",
		columns: sysDeviceKindColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysDeviceKindDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysDeviceKindDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysDeviceKindDao) Columns() SysDeviceKindColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysDeviceKindDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysDeviceKindDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysDeviceKindDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
