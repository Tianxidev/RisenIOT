package logger

import (
	"fmt"
	"time"
)

// LogWriter 日志写入器接口
type LogWriter interface {
	Write(level string, data interface{}) error
}

// Logger 日志器
type Logger struct {
	writerList []LogWriter
}

// NewLogger 创建日志器
func NewLogger() *Logger {
	return &Logger{}
}

// RegisterWriter 注册日志写入器
func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writerList = append(l.writerList, writer)
}

// INFO 记录 INFO 级别的日志
func (l *Logger) INFO(format string, a ...any) {
	for _, writer := range l.writerList {
		writer.Write("INFO", fmt.Sprintf(format, a...))
	}
}

// DEBUG 记录 DEBUG 级别的日志
func (l *Logger) DEBUG(format string, a ...any) {
	for _, writer := range l.writerList {
		writer.Write("DEBUG", fmt.Sprintf(format, a...))
	}
}

// WARNING 记录 WARNING 级别的日志
func (l *Logger) WARNING(format string, a ...any) {
	for _, writer := range l.writerList {
		writer.Write("WARNING", fmt.Sprintf(format, a...))
	}
}

// ERROR 记录 ERROR 级别的日志
func (l *Logger) ERROR(format string, a ...any) {
	for _, writer := range l.writerList {
		writer.Write("ERROR", fmt.Sprintf(format, a...))
	}
}

// LOG 记录 自定义级别的日志
func (l *Logger) LOG(level string, format string, a ...any) {
	for _, writer := range l.writerList {
		writer.Write(level, fmt.Sprintf(format, a...))
	}
}

// CreateLogger 创建日志器的实例
func CreateLogger() *Logger {
	l := NewLogger()

	Y := time.Now().Format("2006")
	M := time.Now().Format("01")
	D := time.Now().Format("02")

	// 创建命令行写入器
	cw := newConsoleWriter()
	l.RegisterWriter(cw)

	// 创建文件写入器
	fw := newFileWriter()
	if err := fw.SetFile(fmt.Sprintf("./logs/%s-%s-%s.log", Y, M, D)); err != nil {
		fmt.Println(err)
	}
	l.RegisterWriter(fw)
	return l
}
