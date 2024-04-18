// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysJob is the golang structure for table sys_job.
type SysJob struct {
	JobId          int    `json:"jobId"          orm:"job_id"          ` //
	JobName        string `json:"jobName"        orm:"job_name"        ` //
	JobParams      string `json:"jobParams"      orm:"job_params"      ` //
	JobGroup       string `json:"jobGroup"       orm:"job_group"       ` //
	InvokeTarget   string `json:"invokeTarget"   orm:"invoke_target"   ` //
	CronExpression string `json:"cronExpression" orm:"cron_expression" ` //
	MisfirePolicy  int    `json:"misfirePolicy"  orm:"misfire_policy"  ` //
	Concurrent     int    `json:"concurrent"     orm:"concurrent"      ` //
	Status         int    `json:"status"         orm:"status"          ` //
	CreateBy       int    `json:"createBy"       orm:"create_by"       ` //
	UpdateBy       int    `json:"updateBy"       orm:"update_by"       ` //
	Remark         string `json:"remark"         orm:"remark"          ` //
	CreatedAt      string `json:"createdAt"      orm:"created_at"      ` //
	UpdatedAt      string `json:"updatedAt"      orm:"updated_at"      ` //
	DeletedAt      string `json:"deletedAt"      orm:"deleted_at"      ` //
}
