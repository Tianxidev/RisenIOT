// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysJobDao is the data access object for table sys_job.
type SysJobDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns SysJobColumns // columns contains all the column names of Table for convenient usage.
}

// SysJobColumns defines and stores column names for table sys_job.
type SysJobColumns struct {
	JobId          string //
	JobName        string //
	JobParams      string //
	JobGroup       string //
	InvokeTarget   string //
	CronExpression string //
	MisfirePolicy  string //
	Concurrent     string //
	Status         string //
	CreateBy       string //
	UpdateBy       string //
	Remark         string //
	CreatedAt      string //
	UpdatedAt      string //
	DeletedAt      string //
}

// sysJobColumns holds the columns for table sys_job.
var sysJobColumns = SysJobColumns{
	JobId:          "job_id",
	JobName:        "job_name",
	JobParams:      "job_params",
	JobGroup:       "job_group",
	InvokeTarget:   "invoke_target",
	CronExpression: "cron_expression",
	MisfirePolicy:  "misfire_policy",
	Concurrent:     "concurrent",
	Status:         "status",
	CreateBy:       "create_by",
	UpdateBy:       "update_by",
	Remark:         "remark",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewSysJobDao creates and returns a new DAO object for table data access.
func NewSysJobDao() *SysJobDao {
	return &SysJobDao{
		group:   "default",
		table:   "sys_job",
		columns: sysJobColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysJobDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysJobDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysJobDao) Columns() SysJobColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysJobDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysJobDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysJobDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
