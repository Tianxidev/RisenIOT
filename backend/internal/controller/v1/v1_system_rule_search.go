package v1

import (
	"backend/internal/model"
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) RuleSearch(ctx context.Context, req *system.RuleSearchReq) (res *system.RuleSearchRes, err error) {
	var list []*model.SysAuthRuleInfoRes
	res = &system.RuleSearchRes{
		Rules: make([]*model.SysAuthRuleTreeRes, 0),
	}
	list, err = service.SysAuthRule().GetMenuListSearch(ctx, req)
	if req.Title != "" || req.Component != "" {
		for _, menu := range list {
			res.Rules = append(res.Rules, &model.SysAuthRuleTreeRes{
				SysAuthRuleInfoRes: menu,
			})
		}
	} else {
		res.Rules = service.SysAuthRule().GetMenuListTree(0, list)
	}
	return
}
