package emqx

import (
	"RisenIOT/backend/internal/env"
	"RisenIOT/backend/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Emqx struct {
}

// EmqxConfig Emqx配置结构体
type EmqxConfig struct {
	Enable            string // Emqx 服务器是否启用
	Panel             string // Emqx 面板地址
	PanelApiKey       string // Emqx 面板ApiKey
	PanelApiSecretKey string // Emqx 面板SecretKey
}

type MqttDeviceList struct {
	Data []MqttDevice `json:"data"`
}

// MqttDevice MQTT设备信息结构体
type MqttDevice struct {
	DeviceId    string `json:"clientid"`     // 设备ID
	IpAddress   string `json:"ip_address"`   // 设备IP地址
	Username    string `json:"username"`     // 设备用户名
	ConnectedAt string `json:"connected_at"` // 设备连接时间
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

// getAuthKey 获取认证密钥
func (d Emqx) getAuthKey() string {
	return base64.StdEncoding.EncodeToString([]byte(d.loadConfig().PanelApiKey + ":" + d.loadConfig().PanelApiSecretKey))
}

// SendDataToTopic 发送消息到指定主题
func (d Emqx) SendDataToTopic(Data string, DataType string, Topic string, qos int) error {

	if Data == "" || Topic == "" {
		return nil
	}

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

	url := d.loadConfig().Panel + "/api/v5/publish"
	method := "POST"

	payload := strings.NewReader(string(payloadBody))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+d.getAuthKey())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// GetDeviceList 获取设备列表
func (d Emqx) GetDeviceList() (*[]MqttDevice, error) {

	url := d.loadConfig().Panel + "/api/v5/clients"

	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Basic " + d.getAuthKey(),
	}

	body, err := utils.HttpGet(url, header)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 反序列化
	var mqttDeviceList MqttDeviceList
	err = json.Unmarshal([]byte(body), &mqttDeviceList)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &mqttDeviceList.Data, nil

}
