// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/api/v1/system"
	"backend/internal/model"
	"backend/internal/model/entity"
	"context"
)

type (
	ISysAuthRule interface {
		GetMenuListSearch(ctx context.Context, req *system.RuleSearchReq) (res []*model.SysAuthRuleInfoRes, err error)
		// GetIsMenuList 获取isMenu=0|1
		GetIsMenuList(ctx context.Context) ([]*model.SysAuthRuleInfoRes, error)
		// GetMenuList 获取所有菜单
		GetMenuList(ctx context.Context) (list []*model.SysAuthRuleInfoRes, err error)
		// GetIsButtonList 获取所有按钮isMenu=2 菜单列表
		GetIsButtonList(ctx context.Context) ([]*model.SysAuthRuleInfoRes, error)
		// Add 添加菜单
		Add(ctx context.Context, req *system.RuleAddReq) (err error)
		// BindRoleRule 绑定角色权限
		BindRoleRule(ctx context.Context, ruleId interface{}, roleIds []uint) (err error)
		Get(ctx context.Context, id uint) (rule *entity.SysAuthRule, err error)
		GetMenuRoles(ctx context.Context, id uint) (roleIds []uint, err error)
		Update(ctx context.Context, req *system.RuleUpdateReq) (err error)
		UpdateRoleRule(ctx context.Context, ruleId uint, roleIds []uint) (err error)
		GetMenuListTree(pid uint, list []*model.SysAuthRuleInfoRes) []*model.SysAuthRuleTreeRes
		// DeleteMenuByIds 删除菜单
		DeleteMenuByIds(ctx context.Context, ids []int) (err error)
		FindSonByParentId(list []*model.SysAuthRuleInfoRes, pid uint) []*model.SysAuthRuleInfoRes
	}
)

var (
	localSysAuthRule ISysAuthRule
)

func SysAuthRule() ISysAuthRule {
	if localSysAuthRule == nil {
		panic("implement not found for interface ISysAuthRule, forgot register?")
	}
	return localSysAuthRule
}

func RegisterSysAuthRule(i ISysAuthRule) {
	localSysAuthRule = i
}
