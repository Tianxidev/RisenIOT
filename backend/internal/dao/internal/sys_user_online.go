// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserOnlineDao is the data access object for table sys_user_online.
type SysUserOnlineDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SysUserOnlineColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserOnlineColumns defines and stores column names for table sys_user_online.
type SysUserOnlineColumns struct {
	Id         string //
	Uuid       string // 用户标识
	Token      string // 用户token
	CreateTime string // 登录时间
	UserName   string // 用户名
	UserId     string // 用户ID
	Ip         string // 登录ip
	Explorer   string // 浏览器
	Os         string // 操作系统
}

// sysUserOnlineColumns holds the columns for table sys_user_online.
var sysUserOnlineColumns = SysUserOnlineColumns{
	Id:         "id",
	Uuid:       "uuid",
	Token:      "token",
	CreateTime: "create_time",
	UserName:   "user_name",
	UserId:     "user_id",
	Ip:         "ip",
	Explorer:   "explorer",
	Os:         "os",
}

// NewSysUserOnlineDao creates and returns a new DAO object for table data access.
func NewSysUserOnlineDao() *SysUserOnlineDao {
	return &SysUserOnlineDao{
		group:   "default",
		table:   "sys_user_online",
		columns: sysUserOnlineColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysUserOnlineDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysUserOnlineDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysUserOnlineDao) Columns() SysUserOnlineColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysUserOnlineDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysUserOnlineDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysUserOnlineDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
