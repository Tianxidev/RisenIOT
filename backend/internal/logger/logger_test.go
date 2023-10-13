package logger

import "testing"

// TestLogger_DEBUG 测试 DEBUG 级别的日志
func TestLogger_DEBUG(t *testing.T) {
	log := CreateLogger()
	log.DEBUG("测试 DEBUG 级别的日志")
}

// TestLogger_INFO 测试 INFO 级别的日志
func TestLogger_INFO(t *testing.T) {
	log := CreateLogger()
	log.INFO("测试 INFO 级别的日志")
}

// TestLogger_WARN 测试 WARN 级别的日志
func TestLogger_WARN(t *testing.T) {
	log := CreateLogger()
	log.WARNING("测试 WARN 级别的日志")
}

// TestLogger_ERROR 测试 ERROR 级别的日志
func TestLogger_ERROR(t *testing.T) {
	log := CreateLogger()
	log.ERROR("测试 ERROR 级别的日志")
}
