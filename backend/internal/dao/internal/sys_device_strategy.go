// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDeviceStrategyDao is the data access object for table sys_device_strategy.
type SysDeviceStrategyDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns SysDeviceStrategyColumns // columns contains all the column names of Table for convenient usage.
}

// SysDeviceStrategyColumns defines and stores column names for table sys_device_strategy.
type SysDeviceStrategyColumns struct {
	Id       string // 策略ID
	Name     string // 策略名称
	Type     string // 策略类型
	Cron     string // Cron时间
	DeviceId string // 设备ID
	AuthorId string // 创建用户ID
	Action   string // 策略动作
}

// sysDeviceStrategyColumns holds the columns for table sys_device_strategy.
var sysDeviceStrategyColumns = SysDeviceStrategyColumns{
	Id:       "id",
	Name:     "name",
	Type:     "type",
	Cron:     "cron",
	DeviceId: "device_id",
	AuthorId: "author_id",
	Action:   "action",
}

// NewSysDeviceStrategyDao creates and returns a new DAO object for table data access.
func NewSysDeviceStrategyDao() *SysDeviceStrategyDao {
	return &SysDeviceStrategyDao{
		group:   "default",
		table:   "sys_device_strategy",
		columns: sysDeviceStrategyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysDeviceStrategyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysDeviceStrategyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysDeviceStrategyDao) Columns() SysDeviceStrategyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysDeviceStrategyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysDeviceStrategyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysDeviceStrategyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
