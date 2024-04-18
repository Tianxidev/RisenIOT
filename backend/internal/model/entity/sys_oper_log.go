// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOperLog is the golang structure for table sys_oper_log.
type SysOperLog struct {
	OperId        uint64      `json:"operId"        orm:"oper_id"        ` // 日志主键
	Title         string      `json:"title"         orm:"title"          ` // 模块标题
	BusinessType  int         `json:"businessType"  orm:"business_type"  ` // 业务类型（0其它 1新增 2修改 3删除）
	Method        string      `json:"method"        orm:"method"         ` // 方法名称
	RequestMethod string      `json:"requestMethod" orm:"request_method" ` // 请求方式
	OperatorType  int         `json:"operatorType"  orm:"operator_type"  ` // 操作类别（0其它 1后台用户 2手机端用户）
	OperName      string      `json:"operName"      orm:"oper_name"      ` // 操作人员
	DeptName      string      `json:"deptName"      orm:"dept_name"      ` // 部门名称
	OperUrl       string      `json:"operUrl"       orm:"oper_url"       ` // 请求URL
	OperIp        string      `json:"operIp"        orm:"oper_ip"        ` // 主机地址
	OperLocation  string      `json:"operLocation"  orm:"oper_location"  ` // 操作地点
	OperParam     string      `json:"operParam"     orm:"oper_param"     ` // 请求参数
	ErrorMsg      string      `json:"errorMsg"      orm:"error_msg"      ` // 错误消息
	OperTime      *gtime.Time `json:"operTime"      orm:"oper_time"      ` // 操作时间
}
