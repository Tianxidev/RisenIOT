// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysRoleDept is the golang structure for table sys_role_dept.
type SysRoleDept struct {
	RoleId int64 `json:"roleId" orm:"role_id" ` // 角色ID
	DeptId int64 `json:"deptId" orm:"dept_id" ` // 部门ID
}
