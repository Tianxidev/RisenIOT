// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/api/v1/system"
	"backend/internal/model/entity"
	"context"
)

type (
	ISysConfig interface {
		// List 系统参数列表
		List(ctx context.Context, req *system.ConfigSearchReq) (res *system.ConfigSearchRes, err error)
		Add(ctx context.Context, req *system.ConfigAddReq, userId uint64) (err error)
		// CheckConfigKeyUnique 验证参数键名是否存在
		CheckConfigKeyUnique(ctx context.Context, configKey string, configId ...int64) (err error)
		// Get 获取系统参数
		Get(ctx context.Context, id int) (res *system.ConfigGetRes, err error)
		// Edit 修改系统参数
		Edit(ctx context.Context, req *system.ConfigEditReq, userId uint64) (err error)
		// Delete 删除系统参数
		Delete(ctx context.Context, ids []int) (err error)
		// GetConfigByKey 通过key获取参数（从缓存获取）
		GetConfigByKey(ctx context.Context, key string) (config *entity.SysConfig, err error)
		// GetByKey 通过key获取参数（从数据库获取）
		GetByKey(ctx context.Context, key string) (config *entity.SysConfig, err error)
	}
)

var (
	localSysConfig ISysConfig
)

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}
