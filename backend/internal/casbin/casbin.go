package casbin

import (
	"RisenIOT/backend/global"
	"fmt"
	"github.com/casbin/casbin/v2"
	"os"
)

// SetupCasbinEnforcer 初始化casbin
func SetupCasbinEnforcer() error {
	dir, _ := os.Getwd()
	modelPath := dir + "/rbac_model.conf"
	csvPath := dir + "/rbac.csv.example"
	global.Logger.INFO("modelPath:" + modelPath)
	global.Logger.INFO("csvPath:" + csvPath)
	var err error
	global.CasbinEnforcer, err = casbin.NewEnforcer(modelPath, csvPath)
	if err != nil {
		global.Logger.ERROR(fmt.Sprintf("启动 Casbin 失败: %v", err))
		return err
	} else {
		global.CasbinEnforcer.EnableLog(true)
	}
	global.Logger.INFO("访问控制模块启用完成")
	return nil
}
