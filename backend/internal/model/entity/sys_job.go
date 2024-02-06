// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysJob is the golang structure for table sys_job.
type SysJob struct {
	JobId          int    `json:"jobId"          ` //
	JobName        string `json:"jobName"        ` //
	JobParams      string `json:"jobParams"      ` //
	JobGroup       string `json:"jobGroup"       ` //
	InvokeTarget   string `json:"invokeTarget"   ` //
	CronExpression string `json:"cronExpression" ` //
	MisfirePolicy  int    `json:"misfirePolicy"  ` //
	Concurrent     int    `json:"concurrent"     ` //
	Status         int    `json:"status"         ` //
	CreateBy       int    `json:"createBy"       ` //
	UpdateBy       int    `json:"updateBy"       ` //
	Remark         string `json:"remark"         ` //
	CreatedAt      string `json:"createdAt"      ` //
	UpdatedAt      string `json:"updatedAt"      ` //
	DeletedAt      string `json:"deletedAt"      ` //
}
