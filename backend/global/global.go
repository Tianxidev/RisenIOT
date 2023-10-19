package global

import (
	"RisenIOT/backend/internal/device"
	"RisenIOT/backend/internal/logger"
	"github.com/casbin/casbin/v2"
)

var (
	SysName         = "RisenIOT"
	SysVersion      = "0.0.1"
	Logger          *logger.Logger
	Device          *device.Device
	DatabaseSetting *databaseSetting
	CasbinEnforcer  *casbin.Enforcer
)

// databaseSetting 数据库配置
type databaseSetting struct {
	DBType    string // 数据库类型
	Host      string // 数据库地址
	DBName    string // 数据库名
	Username  string // 数据库用户名
	Password  string // 数据库密码
	Charset   string // 数据库编码
	ParseTime bool   // 数据库解析时间
}
