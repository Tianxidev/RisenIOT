package response

import "github.com/gin-gonic/gin"

// Error 错误响应
func Error(context *gin.Context, code int, msg string) {
	context.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
	})
}

// Success 成功响应
func Success(context *gin.Context, msg string, data interface{}) {

	if data == nil {
		context.JSON(200, gin.H{
			"code": 200,
			"msg":  msg,
		})
		return
	}

	context.JSON(200, gin.H{
		"code": 200,
		"msg":  msg,
		"data": data,
	})

}
