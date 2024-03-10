// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDeviceCategoryDao is the data access object for table sys_device_category.
type SysDeviceCategoryDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns SysDeviceCategoryColumns // columns contains all the column names of Table for convenient usage.
}

// SysDeviceCategoryColumns defines and stores column names for table sys_device_category.
type SysDeviceCategoryColumns struct {
	Id        string //
	KindId    string // 产品类型ID
	Name      string //
	Mark      string //
	DataType  string // 数据类型
	LogicType string //
	Unit      string // 单位
	Ratio     string // 比例
	Format    string //
	HomeShow  string //
	Remark    string //
	CreatedAt string //
	UpdatedAt string //
	CreateBy  string // 创建人
}

// sysDeviceCategoryColumns holds the columns for table sys_device_category.
var sysDeviceCategoryColumns = SysDeviceCategoryColumns{
	Id:        "id",
	KindId:    "kind_id",
	Name:      "name",
	Mark:      "mark",
	DataType:  "data_type",
	LogicType: "logic_type",
	Unit:      "unit",
	Ratio:     "ratio",
	Format:    "format",
	HomeShow:  "home_show",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	CreateBy:  "create_by",
}

// NewSysDeviceCategoryDao creates and returns a new DAO object for table data access.
func NewSysDeviceCategoryDao() *SysDeviceCategoryDao {
	return &SysDeviceCategoryDao{
		group:   "default",
		table:   "sys_device_category",
		columns: sysDeviceCategoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysDeviceCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysDeviceCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysDeviceCategoryDao) Columns() SysDeviceCategoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysDeviceCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysDeviceCategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysDeviceCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
