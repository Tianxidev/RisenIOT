package env

import "testing"

// LoadEnv_DEBUG 测试加载环境变量
func Test_LoadEnv_DEBUG(t *testing.T) {
	t.Log(GetEnv("APP_WEB_PORT"))
}
