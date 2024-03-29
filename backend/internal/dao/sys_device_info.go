// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"backend/internal/dao/internal"
)

// internalSysDeviceInfoDao is internal type for wrapping internal DAO implements.
type internalSysDeviceInfoDao = *internal.SysDeviceInfoDao

// sysDeviceInfoDao is the data access object for table sys_device_info.
// You can define custom methods on it to extend its functionality as you wish.
type sysDeviceInfoDao struct {
	internalSysDeviceInfoDao
}

var (
	// SysDeviceInfo is globally public accessible object for table sys_device_info operations.
	SysDeviceInfo = sysDeviceInfoDao{
		internal.NewSysDeviceInfoDao(),
	}
)

// Fill with you ideas below.
