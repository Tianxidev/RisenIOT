package system

import "github.com/gogf/gf/v2/frame/g"

// PushTextToPushDeerReq 消息推送测试请求
type PushTextToPushDeerReq struct {
	g.Meta `path:"/system/push/push_deer/text" tags:"PushDeer消息推送" method:"get" summary:"文本消息推送"`
	Key    string `p:"key"`  // 消息推送Key
	Text   string `p:"text"` // 消息内容
}

// PushTextToPushDeerRes 消息推送测试响应
type PushTextToPushDeerRes struct {
	g.Meta `mime:"application/json"`
}

// PushConfigToPushDeerReq 保存 PushDeer 配置请求
type PushConfigToPushDeerReq struct {
	g.Meta `path:"/system/push/push_deer/config" tags:"PushDeer消息推送" method:"post" summary:"保存推送配置"`
	Key    string `p:"key"` // 消息推送Key
}

// PushConfigToPushDeerRes 保存 PushDeer 配置响应
type PushConfigToPushDeerRes struct {
	g.Meta `mime:"application/json"`
}

// PushQueryConfigToPushDeerReq 查询 PushDeer 配置请求
type PushQueryConfigToPushDeerReq struct {
	g.Meta `path:"/system/push/push_deer/config" tags:"PushDeer消息推送" method:"get" summary:"查询推送配置"`
}

// PushQueryConfigToPushDeerRes 查询 PushDeer 配置响应
type PushQueryConfigToPushDeerRes struct {
	g.Meta `mime:"application/json"`
	Key    string `json:"key"` // 消息推送Key
}
