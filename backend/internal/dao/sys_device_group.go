// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"backend/internal/dao/internal"
)

// internalSysDeviceGroupDao is internal type for wrapping internal DAO implements.
type internalSysDeviceGroupDao = *internal.SysDeviceGroupDao

// sysDeviceGroupDao is the data access object for table sys_device_group.
// You can define custom methods on it to extend its functionality as you wish.
type sysDeviceGroupDao struct {
	internalSysDeviceGroupDao
}

var (
	// SysDeviceGroup is globally public accessible object for table sys_device_group operations.
	SysDeviceGroup = sysDeviceGroupDao{
		internal.NewSysDeviceGroupDao(),
	}
)

// Fill with you ideas below.
