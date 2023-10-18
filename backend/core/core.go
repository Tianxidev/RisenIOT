package core

import (
	"RisenIOT/backend/config"
	"RisenIOT/backend/global"
	"RisenIOT/backend/internal/emqx"
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
	global.Logger = logger.CreateLogger()
	global.Logger.INFO("初始化日志模块完成")

	config.SysName, _ = env.GetEnv("APP_NAME")
	global.Logger.INFO("系统名称: " + config.SysName)

	global.Logger.INFO("加载系统配置完成")

	global.Emqx = emqx.CreateEmqx()
	if config.EmqxEnable == "true" {
		global.Logger.INFO("EMQX 服务器已启用")
	}

}

// Start 核心启动函数
func Start() {

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
