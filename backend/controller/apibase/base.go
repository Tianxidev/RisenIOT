package apibase

import (
	"RisenIOT/backend/common/global"
	"RisenIOT/backend/controller/apiresponse"
	"RisenIOT/backend/pkg/logger"
	"RisenIOT/backend/pkg/mycasbin"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type Controller struct {
}

var upgrader = websocket.Upgrader{

	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

// GetVersion 获取系统版本信息
func (e *Controller) GetVersion(context *gin.Context) {
	logger.GlobalLogger.INFO("请求获取系统版本信息")
	context.JSON(200, gin.H{
		"code":          200,
		"SystemName":    global.SysName,
		"SystemVersion": global.SysVersion,
	})
}

// CasbinReload 权限管理重启
func (e *Controller) CasbinReload(context *gin.Context) {
	_, err := mycasbin.LoadPolicy()
	if err != nil {
		apiresponse.Error(context, 500, "加载权限管理模块异常: %v", err)
		return
	}
	apiresponse.Success(context, nil, "权限管理模块加载成功")
}

// ConsoleLogWS 控制台日志WebSocket
func (e *Controller) ConsoleLogWS(context *gin.Context) {

	// 升级get请求为webSocket协议
	wsConn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer wsConn.Close()

	// 判断客户端 5 秒内是否发送了 token
	wsConn.SetReadDeadline(time.Now().Add(5 * time.Second))
	_, b, err := wsConn.ReadMessage()
	if err != nil {
		logger.GlobalLogger.ERROR("读取客户端信息失败: %v", err)
		return
	}
	wsConn.SetReadDeadline(time.Time{})
	token := string(b)
	if token != "WsTerminal" {
		wsConn.WriteMessage(1, []byte("无法识别客户端"))
		wsConn.Close()
		return
	}

	logger.GlobalLogger.INFO("已连接到 Web 日志控制台:" + wsConn.RemoteAddr().String())

	wsConn.WriteMessage(1, []byte("欢迎访问 Web 日志控制台!"))
	wsConn.Close()
}
