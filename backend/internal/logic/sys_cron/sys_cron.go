package sysCron

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
	"time"
)

type sSysCron struct {
}

func New() *sSysCron {
	return &sSysCron{}
}

func init() {
	service.RegisterSysCron(New())
}

// StartBaseTask 启动基础任务
func (s *sSysCron) StartBaseTask(ctx context.Context) {

	// 设备状态检查
	_, err := gcron.AddSingleton(ctx, "* * * * * *", func(ctx context.Context) {
		//g.Log().Info(ctx, "设备保活任务开始")
		service.DeviceInfo().KeepAlive(ctx, gtime.Now())
		//g.Log().Info(ctx, "设备保活任务结束")
		time.Sleep(2 * time.Second)
	})
	liberr.ErrIsNil(ctx, err, "添加任务失败")
}
