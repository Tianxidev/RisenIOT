package emqx

import (
	"RisenIOT/backend/config"
	"RisenIOT/backend/internal/env"
	"errors"
)

type EmqxApi interface {
	sendMsgToTopic(msg string, topic string, qos int) error
}

type Emqx struct {
	emqxList []EmqxApi
}

// NewEmqx 创建Emqx
func NewEmqx() *Emqx {
	return &Emqx{}
}

// LoadConfig 加载配置
func (e *Emqx) LoadConfig() {
	config.EmqxEnable, _ = env.GetEnv("EMQX_ENABLE")
	config.EmqxPanel, _ = env.GetEnv("EMQX_PANEL")
	config.EmqxPanelApiKey, _ = env.GetEnv("EMQX_PANEL_API_KEY")
	config.EmqxPanelApiSecretKey, _ = env.GetEnv("EMQX_PANEL_API_SECRET_KEY")
}

// Register 注册Emqx实例到列表
func (e *Emqx) Register(ea EmqxApi) {
	e.emqxList = append(e.emqxList, ea)
}

// CreateEmqx 创建Emqx实例
func CreateEmqx() *Emqx {
	ne := NewEmqx()
	ne.LoadConfig()
	ndc := newDeviceCmd()
	ne.Register(ndc)
	return ne
}

// DeviceCmdPush 设备命令推送
func (e *Emqx) DeviceCmdPush(cmd string, device string, qos int) error {

	// 推送检查
	if config.EmqxEnable != "true" {
		// 返回错误
		return errors.New("未启用EMQX选项")
	}

	for _, emqx := range e.emqxList {
		err := emqx.sendMsgToTopic(cmd, device, qos)
		if err != nil {
			return err
		}
	}
	return nil
}
