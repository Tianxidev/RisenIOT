package core

import (
	"RisenIOT/backend/internal/logger"
	"RisenIOT/backend/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// Start 核心启动函数
func Start() {

	// 初始化日志器
	log := logger.CreateLogger()

	// 设置 Gin 的运行模式
	gin.SetMode(gin.ReleaseMode)

	// 创建 Gin 引擎
	engine := gin.New()

	// 日志中间件配置
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
			log.LOG("[GIN]", data)
			return data
		},
	}

	engine.Use(gin.LoggerWithConfig(conf))

	PrivateGroup := engine.Group("/api/v1")

	systemRouter := router.GroupApp
	systemRouter.InitBaseRouter(PrivateGroup)

	log.INFO("服务器启动成功")

	// 开启服务器监听
	err := engine.Run(":8080")
	if err != nil {
		return
	}
}
