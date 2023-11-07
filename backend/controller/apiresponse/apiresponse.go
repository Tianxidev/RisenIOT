package apiresponse

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Error 错误响应
func Error(context *gin.Context, code int, formatMsg string, a ...any) {
	context.JSON(200, gin.H{
		"code": code,
		"msg":  fmt.Sprintf(formatMsg, a...),
	})
}

// Success 成功响应
func Success(context *gin.Context, data interface{}, formatMsg string, a ...any) {

	if data == nil {
		context.JSON(200, gin.H{
			"code": 200,
			"msg":  fmt.Sprintf(formatMsg, a...),
		})
		return
	}

	context.JSON(200, gin.H{
		"code": 200,
		"msg":  fmt.Sprintf(formatMsg, a...),
		"data": data,
	})

}
