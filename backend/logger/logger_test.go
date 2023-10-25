package logger

import "testing"

// Test_Logger_DEBUG 测试 DEBUG 级别的日志
func Test_Logger_DEBUG(t *testing.T) {
	InitGlobalLogger()
	GlobalLogger.DEBUG("测试 DEBUG 级别的日志")
}

// Test_Logger_INFO 测试 INFO 级别的日志
func Test_Logger_INFO(t *testing.T) {
	InitGlobalLogger()
	GlobalLogger.INFO("测试 INFO 级别的日志")
}

// Test_Logger_WARN 测试 WARN 级别的日志
func Test_Logger_WARN(t *testing.T) {
	InitGlobalLogger()
	GlobalLogger.WARNING("测试 WARN 级别的日志")
}

// Test_Logger_ERROR 测试 ERROR 级别的日志
func Test_Logger_ERROR(t *testing.T) {
	InitGlobalLogger()
	GlobalLogger.ERROR("测试 ERROR 级别的日志")
}
