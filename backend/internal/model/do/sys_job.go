// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysJob is the golang structure of table sys_job for DAO operations like Where/Data.
type SysJob struct {
	g.Meta         `orm:"table:sys_job, do:true"`
	JobId          interface{} //
	JobName        interface{} //
	JobParams      interface{} //
	JobGroup       interface{} //
	InvokeTarget   interface{} //
	CronExpression interface{} //
	MisfirePolicy  interface{} //
	Concurrent     interface{} //
	Status         interface{} //
	CreateBy       interface{} //
	UpdateBy       interface{} //
	Remark         interface{} //
	CreatedAt      interface{} //
	UpdatedAt      interface{} //
	DeletedAt      interface{} //
}
