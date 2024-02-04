// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"backend/internal/dao/internal"
)

// internalSysDeptDao is internal type for wrapping internal DAO implements.
type internalSysDeptDao = *internal.SysDeptDao

// sysDeptDao is the data access object for table sys_dept.
// You can define custom methods on it to extend its functionality as you wish.
type sysDeptDao struct {
	internalSysDeptDao
}

var (
	// SysDept is globally public accessible object for table sys_dept operations.
	SysDept = sysDeptDao{
		internal.NewSysDeptDao(),
	}
)

// Fill with you ideas below.
