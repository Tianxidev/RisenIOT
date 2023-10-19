package emqx

import (
	"RisenIOT/backend/internal/env"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Emqx struct {
}

type EmqxConfig struct {
	Enable            string // Emqx 服务器是否启用
	Panel             string // Emqx 面板地址
	PanelApiKey       string // Emqx 面板ApiKey
	PanelApiSecretKey string // Emqx 面板SecretKey
}

// Create 创建设备命令实例
func Create() *Emqx {
	return &Emqx{}
}

// loadConfig 加载配置
func (d *Emqx) loadConfig() EmqxConfig {
	var emqxConfig EmqxConfig
	emqxConfig.Enable, _ = env.GetEnv("EMQX_ENABLE")
	emqxConfig.Panel, _ = env.GetEnv("EMQX_PANEL")
	emqxConfig.PanelApiKey, _ = env.GetEnv("EMQX_PANEL_API_KEY")
	emqxConfig.PanelApiSecretKey, _ = env.GetEnv("EMQX_PANEL_API_SECRET_KEY")
	return emqxConfig
}

// SendDataToTopic 发送消息到指定主题
func (d Emqx) SendDataToTopic(Data string, DataType string, Topic string, qos int) error {

	// 处理消息载荷类型, 如果不为 base64 则以字符串的形式传输
	if DataType != "base64" {
		DataType = "plain"
	}

	payloadBodyOrigin := map[string]interface{}{"payload_encoding": DataType, "topic": Topic, "qos": 1, "payload": Data, "retain": false}
	payloadBody, err := json.Marshal(payloadBodyOrigin)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return err
	}

	fmt.Println("发送消息: ", string(payloadBody))

	url := d.loadConfig().Panel + "/api/v5/publish"
	method := "POST"

	// 加密认真信息
	sEnc := base64.StdEncoding.EncodeToString([]byte(d.loadConfig().PanelApiKey + ":" + d.loadConfig().PanelApiSecretKey))

	fmt.Println("认证信息: ", sEnc)
	fmt.Println("接口地址: ", url)

	payload := strings.NewReader(string(payloadBody))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+sEnc)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("EMQX API 响应: " + string(body))
	return nil
}
