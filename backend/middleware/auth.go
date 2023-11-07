package middleware

import (
	"RisenIOT/backend/models"
	"RisenIOT/backend/pkg/logger"
	"RisenIOT/backend/pkg/mycasbin"
	"RisenIOT/backend/pkg/myjwt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Auth 权限中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sub := "null"
		username := "null"

		// 读取url
		obj := c.Request.URL.RequestURI()

		// 获取 header 中的 token
		authorization := c.Request.Header.Get("authorization")

		// 判断是否存在 token
		if authorization != "" {
			var sysUser models.SysUser
			// 解密
			mc, err := myjwt.ParseToken(authorization)
			if err != nil {
				logger.GlobalLogger.ERROR("解析token失败: %v", err)
			}

			sysUser.UserId = int(mc.UserID)

			if sysUser, err = sysUser.GetUserFromId(); err == nil {
				username = sysUser.Username
				sub = strconv.Itoa(sysUser.RoleId)
			} else {
				logger.GlobalLogger.ERROR("获取用户信息失败: %v", err)
			}

		}

		// 获取方法
		act := c.Request.Method

		// 获取 casbin 示例
		e := mycasbin.Casbin()

		// 判断策略是否已经存在
		if ok, _ := e.Enforce(sub, obj, act); ok {
			logger.GlobalLogger.SUCCESS("用户 %s(%s) 访问URL: [%s] %s ", username, sub, act, obj)
			c.Next()
		} else {
			logger.GlobalLogger.INFO("用户 %s 请求权限不足", sub)
			c.JSON(200, gin.H{
				"code": 403,
				"msg":  "权限不足",
			})
			c.Abort()
		}
	}
}
