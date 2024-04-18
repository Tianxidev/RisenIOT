package streaming_media

import "github.com/gogf/gf/v2/frame/g"

// PullFlowForwardingReq 拉流转发请求
type PullFlowForwardingReq struct {
	g.Meta `path:"/streaming_media/pull_flow_forwarding" tags:"流媒体服务" method:"get" summary:"拉流转发"`
	Rtsp   string `json:"rtsp" v:"required|length:1,255#请输入rtsp地址|rtsp地址长度为:min到:max位"`
	Srs    string `json:"srs" v:"required|length:1,255#请输入srs地址|srs地址长度为:min到:max位"`
}

// PullFlowForwardingRes 拉流转发响应
type PullFlowForwardingRes struct {
	Stream string `json:"stream"` // 流名称
	FlvSrc string `json:"flvSrc"` // flv源地址
}
