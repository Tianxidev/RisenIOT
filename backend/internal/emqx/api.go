package emqx

import (
	"RisenIOT/backend/config"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type deviceCmd struct {
}

func (d deviceCmd) sendMsgToTopic(msg string, topic string, qos int) error {

	//payloadBody := `{"payload_encoding": "plain",　"topic": "` + topic + `",　"qos": 1,　"payload": "` + msg + `",　"retain": false}`

	payloadBodyOrigin := map[string]interface{}{"payload_encoding": "plain", "topic": topic, "qos": 1, "payload": msg, "retain": false}
	payloadBody, err := json.Marshal(payloadBodyOrigin)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return err
	}

	fmt.Println("发送消息: ", payloadBody)

	url := config.EmqxPanel + "/api/v5/publish"
	method := "POST"

	// 加密认真信息
	sEnc := base64.StdEncoding.EncodeToString([]byte(config.EmqxPanelApiKey + ":" + config.EmqxPanelApiSecretKey))

	fmt.Println("认证信息: ", sEnc)

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
	fmt.Println(string(body))
	return nil
}

// NewDeviceCmd 创建设备命令实例
func newDeviceCmd() *deviceCmd {
	return &deviceCmd{}
}
