package apibase

import (
	"RisenIOT/backend/controller/apiresponse"
	"RisenIOT/backend/models"
	"RisenIOT/backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

// Login 登录
func (this *Controller) Login(context *gin.Context) {
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
	user, e := loginVals.GetUser()
	if e == nil {
		apiresponse.Success(context, user, msg)
	} else {
		msg = "登录失败"
		code = 1
		apiresponse.Error(context, code, msg)
	}

	logger.GlobalLogger.DEBUG("用户 %s 登录, 结果: %s", username, msg)

	return

}
