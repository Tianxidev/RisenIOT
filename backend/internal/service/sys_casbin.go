// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
)

type (
	IAdapterCasbin interface {
		// CasbinEnforcer 获取adapter单例对象
		CasbinEnforcer(ctx context.Context) (enforcer *casbin.SyncedEnforcer, err error)
		// SavePolicy saves policy to database.
		SavePolicy(model model.Model) (err error)
		// LoadPolicy loads policy from database.
		LoadPolicy(model model.Model) error
		// AddPolicy adds a policy rule to the storage.
		AddPolicy(sec string, ptype string, rule []string) error
		// RemovePolicy removes a policy rule from the storage.
		RemovePolicy(sec string, ptype string, rule []string) error
		AddPolicies(sec string, ptype string, rules [][]string) error
		// RemovePolicies removes policy rules from the storage.
		// This is part of the Auto-Save feature.
		RemovePolicies(sec string, ptype string, rules [][]string) error
		// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
		RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error
	}
)

var (
	localAdapterCasbin IAdapterCasbin
)

func AdapterCasbin() IAdapterCasbin {
	if localAdapterCasbin == nil {
		panic("implement not found for interface IAdapterCasbin, forgot register?")
	}
	return localAdapterCasbin
}

func RegisterAdapterCasbin(i IAdapterCasbin) {
	localAdapterCasbin = i
}
