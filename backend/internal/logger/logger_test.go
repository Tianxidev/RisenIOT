package logger

import "testing"

// Test_Logger_DEBUG 测试 DEBUG 级别的日志
func Test_Logger_DEBUG(t *testing.T) {
	log := CreateLogger()
	log.DEBUG("测试 DEBUG 级别的日志")
}

// Test_Logger_INFO 测试 INFO 级别的日志
func Test_Logger_INFO(t *testing.T) {
	log := CreateLogger()
	log.INFO("测试 INFO 级别的日志")
}

// Test_Logger_WARN 测试 WARN 级别的日志
func Test_Logger_WARN(t *testing.T) {
	log := CreateLogger()
	log.WARNING("测试 WARN 级别的日志")
}

// Test_Logger_ERROR 测试 ERROR 级别的日志
func Test_Logger_ERROR(t *testing.T) {
	log := CreateLogger()
	log.ERROR("测试 ERROR 级别的日志")
}
