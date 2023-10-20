package controller

import (
	"RisenIOT/backend/global"
	"RisenIOT/backend/internal/casbin"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type BaseController struct {
}

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func NewBaseController() *BaseController {
	return &BaseController{}
}

// GetVersion 获取系统版本信息
func (nbc *BaseController) GetVersion(c *gin.Context) {
	global.Logger.INFO("请求获取系统版本信息")
	c.JSON(200, gin.H{
		"code":       200,
		"SysName":    global.SysName,
		"SysVersion": global.SysVersion,
	})
}

// CasbinReload 权限管理重启
func (nbc *BaseController) CasbinReload(c *gin.Context) {
	casbin.SetupCasbinEnforcer()
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

// ConsoleLogWS 控制台日志WebSocket
func (nbc *BaseController) ConsoleLogWS(c *gin.Context) {

	// 升级get请求为webSocket协议
	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer wsConn.Close()

	// 判断客户端
	wsConn.SetReadDeadline(time.Now().Add(5 * time.Second))
	_, b, err := wsConn.ReadMessage()
	if err != nil {
		global.Logger.ERROR("读取客户端信息失败: %v", err)
		return
	}
	wsConn.SetReadDeadline(time.Time{})
	token := string(b)
	if token != "WsTerminal" {
		wsConn.WriteMessage(1, []byte("Invalid client token"))
		wsConn.Close()
		return
	}
	wsConn.WriteMessage(1, []byte("token success"))
	wsConn.Close()

}
