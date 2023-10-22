package emqx

type T struct {
	Username  string `json:"username"`   // 设备用户名
	Timestamp int64  `json:"timestamp"`  // 时间戳
	SockName  string `json:"sockname"`   // 连接节点地址
	Reason    string `json:"reason"`     // 断开原因
	ProtoVer  int    `json:"proto_ver"`  // 协议版本
	ProtoName string `json:"proto_name"` // 协议名称
	PeerName  string `json:"peername"`
	Node      string `json:"node"`
	Metadata  struct {
		RuleId string `json:"rule_id"` // 规则ID
	} `json:"metadata"`
	Event          string `json:"event"`           // 事件
	DisconnectedAt int64  `json:"disconnected_at"` // 断开时间
	DisConnProps   struct {
		UserProperty struct {
		} `json:"User-Property"`
	} `json:"disconn_props"`
	ClientId string `json:"clientid"` // 设备ID
}
