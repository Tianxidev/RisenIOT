package core

import (
	"RisenIOT/backend/casbin"
	"RisenIOT/backend/controller/response"
	"RisenIOT/backend/device"
	"RisenIOT/backend/env"
	"RisenIOT/backend/global"
	"RisenIOT/backend/logger"
	"RisenIOT/backend/middleware"
	"RisenIOT/backend/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
)

func Init() {

	// 初始化日志模块
	logger.InitGlobalLogger()

	// 读取配置文件
	global.SysName, _ = env.GetEnv("APP_NAME")
	logger.GlobalLogger.INFO("欢迎使用" + global.SysName + "系统")
	logger.GlobalLogger.INFO("加载系统配置完成")

	// 打印系统版本号
	logger.GlobalLogger.INFO("系统版本号: " + global.SysVersion)

	// 初始化设备管理模块
	global.Device = device.CreateDevice()
	logger.GlobalLogger.INFO("初始化设备管理模块")

	// casbin 初始化
	err := casbin.SetupCasbinEnforcer()
	if err != nil {
		logger.GlobalLogger.ERROR("初始化权限管理模块异常: %v", err)
	}

}

// Enable 核心启动函数
func Enable() {

	// 配置 gin 为生产模式
	gin.SetMode(gin.ReleaseMode)

	// 初始化 gin
	engine := gin.New()

	// 配置 gin 日志
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
			logger.GlobalLogger.LOG("GIN", data)
			return data
		},
	}

	// 配置 gin 日志中间件
	engine.Use(gin.LoggerWithConfig(conf))

	// 配置 gin 同源策略中间件
	engine.Use(middleware.Cors())

	// 配置 gin 认证中间件
	engine.Use(middleware.Auth(global.CasbinEnforcer))

	// 创建路由组
	v1ApiGroup := engine.Group("/api/v1")

	// 初始化路由
	router.InitRouter(v1ApiGroup)

	// 查询路由列表
	engine.GET("/api/list", func(c *gin.Context) {

		// 创建json map
		json := make(map[string]interface{})

		routers := engine.Routes()
		for _, v := range routers {
			json[v.Path] = v.Method
		}

		// 返回json
		response.Success(c, "获取路由列表成功", json)

	})

	// 读取环境变量中的端口号
	PORT, err := env.GetEnv("APP_WEB_PORT")
	if err != nil {
		logger.GlobalLogger.ERROR("无法读取环境变量中的端口号")
		os.Exit(0)
	}

	// 启动 gin
	logger.GlobalLogger.INFO("Web 服务启动中, 工作端口:" + PORT)
	err = engine.Run(":" + PORT)
	if err != nil {
		log.Fatalf("Web 服务启动失败: %v", err)
		os.Exit(0)
	}
}
