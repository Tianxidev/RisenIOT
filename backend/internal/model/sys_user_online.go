package model

// SysUserOnlineParams 用户在线状态写入参数
type SysUserOnlineParams struct {
	UserAgent string
	Uuid      string
	Token     string
	Username  string
	Ip        string
}
