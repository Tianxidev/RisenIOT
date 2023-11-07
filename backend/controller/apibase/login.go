package apibase

import (
	"RisenIOT/backend/controller/apiresponse"
	"RisenIOT/backend/models"
	"RisenIOT/backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

// Login 登录
func (e *Controller) Login(context *gin.Context) {
	var loginVals models.Login
	var code = 200
	var msg = "登录成功"
	var username = ""

	// 解析请求体数据
	if err := context.ShouldBind(&loginVals); err != nil {
		username = loginVals.Username
		msg = "数据解析失败"
		code = 1
		apiresponse.Error(context, code, msg)
		return
	}

	username = loginVals.Username

	// 获取用户信息
	if _, err := loginVals.GetUser(); err == nil {

		// 生成 jwt token
		token, err := loginVals.GenToken()
		if err != nil {
			return
		}
		apiresponse.Success(context, token, msg)

	} else {
		msg = "登录失败"
		code = 1
		apiresponse.Error(context, code, msg)
	}

	logger.GlobalLogger.DEBUG("用户 %s 登录, 结果: %s", username, msg)

	return

}
