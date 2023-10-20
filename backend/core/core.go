package core

import (
	"RisenIOT/backend/global"
	"RisenIOT/backend/internal/casbin"
	"RisenIOT/backend/internal/device"
	"RisenIOT/backend/internal/env"
	"RisenIOT/backend/internal/logger"
	"RisenIOT/backend/middleware"
	"RisenIOT/backend/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

func Init() {

	// 初始化日志模块
	global.Logger = logger.CreateLogger()

	// 读取配置文件
	global.SysName, _ = env.GetEnv("APP_NAME")
	global.Logger.INFO("欢迎使用" + global.SysName + "系统")
	global.Logger.INFO("加载系统配置完成")

	// 初始化设备管理模块
	global.Device = device.CreateDevice()
	global.Logger.INFO("初始化设备管理模块")

	// casbin 初始化
	err := casbin.SetupCasbinEnforcer()
	if err != nil {
		global.Logger.ERROR("初始化权限管理模块异常: %v", err)
	}

}

// Enable 核心启动函数
func Enable() {

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	conf := gin.LoggerConfig{
		SkipPaths: []string{},
		Output:    ioutil.Discard,
		Formatter: func(p gin.LogFormatterParams) string {
			data := fmt.Sprintf("%d %s %s %s",
				p.StatusCode,
				p.Method,
				p.Path,
				p.ClientIP,
			)
			global.Logger.LOG("GIN", data)
			return data
		},
	}

	engine.Use(gin.LoggerWithConfig(conf))
	engine.Use(middleware.Cors())
	engine.Use(middleware.Auth(global.CasbinEnforcer))

	PrivateGroup := engine.Group("/api/v1")

	systemRouter := router.GroupApp
	systemRouter.InitBaseRouter(PrivateGroup)
	systemRouter.InitDeviceRouter(PrivateGroup)

	PORT, err := env.GetEnv("APP_WEB_PORT")
	if err != nil {
		global.Logger.ERROR("无法读取环境变量中的端口号")
		os.Exit(0)
	}

	global.Logger.INFO("Web 服务启动中, 工作端口:" + PORT)

	engine.Run(":" + PORT)
}
