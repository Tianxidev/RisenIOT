package casbin

import (
	"RisenIOT/backend/common/global"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
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

	//csvPath := "./config/rbac.csv"
	//logger.GlobalLogger.INFO("加载路由控制表:" + csvPath)

	//Apter, err := gormAdapter.NewAdapterByDBUsePrefix(global.Eloquent, "sys_")
	Apter, err := gormAdapter.NewAdapterByDB(global.Eloquent)
	if err != nil {
		panic(err)
	}
	m, err := model.NewModelFromString(modelText)
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewSyncedEnforcer(m, Apter)
	if err != nil {
		panic(err)
	}
	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}
	global.CasbinEnforcer = e
}
