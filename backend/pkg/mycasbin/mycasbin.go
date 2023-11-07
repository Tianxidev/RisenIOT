package mycasbin

import (
	"RisenIOT/backend/common/global"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"log"
)

var Enforcer *casbin.Enforcer

// modelText Casbin模型
var modelText = `
# Request definition 自定义请求的格式
[request_definition]
r = sub, obj, act

# Policy definition 策略定义
[policy_definition]
p = sub, obj, act

# Policy effect 策略生效的规则
[policy_effect]
e = some(where (p.eft == allow))

# Matchers 匹配器
[matchers]
m = (r.sub == p.sub || p.sub == "*") && keyMatch(r.obj,p.obj) && (r.act == p.act || p.act == "*")
`

// Setup 初始化casbin
func Setup() {
	var err error

	// 创建 casbin 适配器
	Apter, err := gormAdapter.NewAdapterByDB(global.Eloquent)
	if err != nil {
		panic(err)
	}

	// 创建模型
	m, err := model.NewModelFromString(modelText)
	if err != nil {
		panic(err)
	}

	// 创建 casbin 实例
	e, err := casbin.NewSyncedEnforcer(m, Apter)
	if err != nil {
		panic(err)
	}

	// 加载策略
	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}

	global.CasbinEnforcer = e

}

// Casbin 获取casbin实例
func Casbin() *casbin.SyncedEnforcer {
	return global.CasbinEnforcer
}

// LoadPolicy 加载策略
func LoadPolicy() (*casbin.SyncedEnforcer, error) {
	if err := global.CasbinEnforcer.LoadPolicy(); err == nil {
		return global.CasbinEnforcer, err
	} else {
		log.Printf("casbin rbac_model or policy init error, message: %v \r\n", err.Error())
		return nil, err
	}
}
