package casbin

import (
	"RisenIOT/backend/logger"
	"github.com/casbin/casbin/v2"
)

var Enforcer *casbin.Enforcer

// SetupCasbinEnforcer 初始化casbin
func SetupCasbinEnforcer() error {
	modelPath := "./config/rbac_model.conf"
	csvPath := "./config/rbac.csv"
	logger.GlobalLogger.INFO("加载访问控制模型:" + modelPath)
	logger.GlobalLogger.INFO("加载路由控制表:" + csvPath)
	var err error
	Enforcer, err = casbin.NewEnforcer(modelPath, csvPath)
	if err != nil {
		logger.GlobalLogger.ERROR("启动 Casbin 失败: %v", err)
		return err
	}
	logger.GlobalLogger.INFO("访问控制模块启用完成")
	return nil
}
