// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RiseniotUserDao is the data access object for table riseniot_user.
type RiseniotUserDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns RiseniotUserColumns // columns contains all the column names of Table for convenient usage.
}

// RiseniotUserColumns defines and stores column names for table riseniot_user.
type RiseniotUserColumns struct {
	Uuid     string // 用户uuid
	Username string // 用户名, 唯一索引
	Password string // 密码
	Nickname string // 昵称
	Role     string // 用户角色
}

// riseniotUserColumns holds the columns for table riseniot_user.
var riseniotUserColumns = RiseniotUserColumns{
	Uuid:     "uuid",
	Username: "username",
	Password: "password",
	Nickname: "nickname",
	Role:     "role",
}

// NewRiseniotUserDao creates and returns a new DAO object for table data access.
func NewRiseniotUserDao() *RiseniotUserDao {
	return &RiseniotUserDao{
		group:   "default",
		table:   "riseniot_user",
		columns: riseniotUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RiseniotUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RiseniotUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RiseniotUserDao) Columns() RiseniotUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RiseniotUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RiseniotUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RiseniotUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
