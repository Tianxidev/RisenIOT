package gateway

import (
	"backend/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtcp"
	"net"
	"strconv"
)

type (
	sGateway struct{}

	// client 客户端, 用于接收各个业务端传入的请求
	client struct {
		conn   net.Conn    // 网络连接对象
		read   chan []byte // 数据读取通道
		write  chan []byte // 数据写入通道
		exit   chan error  // 异常退出通道
		reConn chan bool   // 重连通道
	}

	// user 用户端, 用于接收用户浏览器的连接
	user struct {
		conn   net.Conn    // 网络连接对象
		read   chan []byte // 数据读取通道
		write  chan []byte // 数据写入通道
		exit   chan error  // 异常退出通道
		reConn chan bool   // 重连通道
	}
)

func New() service.IGateway {
	return &sGateway{}
}

func init() {
	service.RegisterGateway(New())
}

// Start 网关启动
func (s *sGateway) Start(ctx context.Context) {
	var (
		defaultPort = 80
		cfg         = g.Cfg()
	)

	// 设置 tcp 端口
	serverPort, _ := cfg.Get(ctx, "gateway.port")
	if serverPort == nil {
		g.Log().Notice(ctx, "读取工作端口失败, 将使用默认端口: "+strconv.Itoa(defaultPort))
	} else {
		g.Log().Info(ctx, "gateway 服务设置端口 ["+strconv.Itoa(serverPort.Int())+"] 成功!")
		defaultPort = serverPort.Int()
	}

	// 启动 tcp 服务
	err := gtcp.NewServer("0.0.0.0:"+strconv.Itoa(defaultPort), func(conn *gtcp.Conn) {
		defer func(conn *gtcp.Conn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)
		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				if err := conn.Send(append([]byte("> "), data...)); err != nil {
					fmt.Println(err)
				}
			}
			if err != nil {
				break
			}
		}
	}).Run()
	if err != nil {
		return
	}
}
