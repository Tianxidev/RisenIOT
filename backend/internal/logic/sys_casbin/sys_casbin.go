package sysCasbin

import (
	"backend/internal/dao"
	"backend/internal/model/entity"
	"backend/internal/service"
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/gogf/gf/v2/frame/g"
	"sync"
)

type sAdapterCasbin struct {
	Enforcer    *casbin.SyncedEnforcer
	EnforcerErr error
	ctx         context.Context
}

func init() {
	service.RegisterAdapterCasbin(New())
}

func New() service.IAdapterCasbin {
	return &sAdapterCasbin{}
}

var (
	once  sync.Once
	en    *casbin.SyncedEnforcer
	enErr error
)

// CasbinEnforcer 获取adapter单例对象
func (_ *sAdapterCasbin) CasbinEnforcer(ctx context.Context) (enforcer *casbin.SyncedEnforcer, err error) {
	ac := newAdapter(ctx)
	enforcer = ac.Enforcer
	err = ac.EnforcerErr
	return
}

// 初始化adapter操作
func newAdapter(ctx context.Context) (a *sAdapterCasbin) {
	a = new(sAdapterCasbin)
	a.ctx = ctx
	once.Do(func() {
		en, enErr = initPolicy(ctx, a)
	})
	if enErr == nil && en != nil {
		en.SetAdapter(a)
	}
	a.Enforcer, a.EnforcerErr = en, enErr
	return
}

func initPolicy(ctx context.Context, a *sAdapterCasbin) (e *casbin.SyncedEnforcer, err error) {
	// Because the DB is empty at first,
	// so we need to load the policy from the file adapter (.CSV) first.
	e, err = casbin.NewSyncedEnforcer(g.Cfg().MustGet(ctx, "casbin.modelFile").String(), a)
	if err != nil {
		g.Log().Error(ctx, "NewSyncedEnforcer->", err)

	}
	return
}

// SavePolicy saves policy to database.
func (a *sAdapterCasbin) SavePolicy(model model.Model) (err error) {
	err = a.dropTable()
	if err != nil {
		return
	}
	err = a.createTable()
	if err != nil {
		return
	}
	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			_, err := dao.CasbinRule.Ctx(a.ctx).Data(line).Insert()
			if err != nil {
				return err
			}
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			_, err := dao.CasbinRule.Ctx(a.ctx).Data(line).Insert()
			if err != nil {
				return err
			}
		}
	}
	return
}

func (a *sAdapterCasbin) dropTable() (err error) {
	return
}

func (a *sAdapterCasbin) createTable() (err error) {
	return
}

// LoadPolicy loads policy from database.
func (a *sAdapterCasbin) LoadPolicy(model model.Model) error {
	var lines []*entity.CasbinRule
	if err := dao.CasbinRule.Ctx(a.ctx).Scan(&lines); err != nil {
		return err
	}
	for _, line := range lines {
		loadPolicyLine(line, model)
	}
	return nil
}

// AddPolicy adds a policy rule to the storage.
func (a *sAdapterCasbin) AddPolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	_, err := dao.CasbinRule.Ctx(a.ctx).Data(line).Insert()
	return err
}

// RemovePolicy removes a policy rule from the storage.
func (a *sAdapterCasbin) RemovePolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	err := rawDelete(a, line)
	return err
}

func (a *sAdapterCasbin) AddPolicies(sec string, ptype string, rules [][]string) error {
	lines := make([]*entity.CasbinRule, len(rules))
	for k, rule := range rules {
		lines[k] = savePolicyLine(ptype, rule)
	}
	_, err := dao.CasbinRule.Ctx(a.ctx).Data(lines).Insert()
	return err
}

// RemovePolicies removes policy rules from the storage.
// This is part of the Auto-Save feature.
func (a *sAdapterCasbin) RemovePolicies(sec string, ptype string, rules [][]string) error {
	for _, rule := range rules {
		err := a.RemovePolicy(sec, ptype, rule)
		if err != nil {
			return err
		}
	}
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *sAdapterCasbin) RemoveFilteredPolicy(sec string, ptype string,
	fieldIndex int, fieldValues ...string) error {
	line := &entity.CasbinRule{}
	line.Ptype = ptype
	if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) {
		line.V0 = fieldValues[0-fieldIndex]
	}
	if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) {
		line.V1 = fieldValues[1-fieldIndex]
	}
	if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) {
		line.V2 = fieldValues[2-fieldIndex]
	}
	if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) {
		line.V3 = fieldValues[3-fieldIndex]
	}
	if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) {
		line.V4 = fieldValues[4-fieldIndex]
	}
	if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) {
		line.V5 = fieldValues[5-fieldIndex]
	}
	err := rawDelete(a, line)
	return err
}

func loadPolicyLine(line *entity.CasbinRule, model model.Model) {
	lineText := line.Ptype
	if line.V0 != "" {
		lineText += ", " + line.V0
	}
	if line.V1 != "" {
		lineText += ", " + line.V1
	}
	if line.V2 != "" {
		lineText += ", " + line.V2
	}
	if line.V3 != "" {
		lineText += ", " + line.V3
	}
	if line.V4 != "" {
		lineText += ", " + line.V4
	}
	if line.V5 != "" {
		lineText += ", " + line.V5
	}
	persist.LoadPolicyLine(lineText, model)
}

func savePolicyLine(ptype string, rule []string) *entity.CasbinRule {
	line := &entity.CasbinRule{}
	line.Ptype = ptype
	if len(rule) > 0 {
		line.V0 = rule[0]
	}
	if len(rule) > 1 {
		line.V1 = rule[1]
	}
	if len(rule) > 2 {
		line.V2 = rule[2]
	}
	if len(rule) > 3 {
		line.V3 = rule[3]
	}
	if len(rule) > 4 {
		line.V4 = rule[4]
	}
	if len(rule) > 5 {
		line.V5 = rule[5]
	}
	return line
}

func rawDelete(a *sAdapterCasbin, line *entity.CasbinRule) error {
	db := dao.CasbinRule.Ctx(a.ctx).Where("ptype = ?", line.Ptype)
	if line.V0 != "" {
		db = db.Where("v0 = ?", line.V0)
	}
	if line.V1 != "" {
		db = db.Where("v1 = ?", line.V1)
	}
	if line.V2 != "" {
		db = db.Where("v2 = ?", line.V2)
	}
	if line.V3 != "" {
		db = db.Where("v3 = ?", line.V3)
	}
	if line.V4 != "" {
		db = db.Where("v4 = ?", line.V4)
	}
	if line.V5 != "" {
		db = db.Where("v5 = ?", line.V5)
	}
	_, err := db.Delete()
	return err
}
