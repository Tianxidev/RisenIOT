package middleware

import (
	"RisenIOT/backend/pkg/logger"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// Auth 权限中间件
func Auth(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		obj := c.Request.URL.RequestURI()
		// 获取方法
		act := c.Request.Method
		sub := "root"

		// 判断策略是否已经存在
		if ok, _ := e.Enforce(sub, obj, act); ok {
			c.Next()
		} else {
			logger.GlobalLogger.INFO("用户请求权限不足")
			c.JSON(200, gin.H{
				"code": 403,
				"msg":  "权限不足",
			})
			c.Abort()
		}
	}
}
