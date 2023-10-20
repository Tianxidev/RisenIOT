package casbin

import (
	"RisenIOT/backend/global"
	"github.com/casbin/casbin/v2"
	"os"
)

// SetupCasbinEnforcer 初始化casbin
func SetupCasbinEnforcer() error {
	dir, _ := os.Getwd()
	modelPath := dir + "\\config\\rbac_model.conf"
	csvPath := dir + "\\config\\rbac.csv"
	global.Logger.INFO("加载访问控制模型:" + modelPath)
	global.Logger.INFO("加载路由控制表:" + csvPath)
	var err error
	global.CasbinEnforcer, err = casbin.NewEnforcer(modelPath, csvPath)
	if err != nil {
		global.Logger.ERROR("启动 Casbin 失败: %v", err)
		return err
	}
	global.Logger.INFO("访问控制模块启用完成")
	return nil
}
