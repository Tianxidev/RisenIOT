package logger

import (
	"errors"
	"fmt"
	"os"
)

// fileWriter 文件写入器
type fileWriter struct {
	file *os.File
}

// consoleWriter 控制台写入器
type consoleWriter struct {
}

// WebSocketLogChan 字符串日志通道
type WebSocketLogChan chan string

// SetFile 设置文件写入器写入的文件名
func (f *fileWriter) SetFile(filename string) (err error) {
	if f.file != nil {
		f.file.Close()
	}

	// 检查文件是否存在，不存在则创建
	_, err = os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		f.file, err = os.Create(filename)
		return err
	} else {
		f.file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666)
		return err
	}
}

// Write 写入文件
func (f *fileWriter) Write(level string, data interface{}) error {
	if f.file == nil {
		return errors.New("写入日志时发现日志文件不存在")
	}
	str := fmt.Sprintf("[%s][%s]: %v\r\n", GetTimeString(), level, data)
	_, err := f.file.Write([]byte(str))
	return err
}

// 实现LogWriter的Write()方法
func (f *consoleWriter) Write(level string, data interface{}) error {
	str := fmt.Sprintf("[%s][%s]: %v\r\n", GetTimeString(), level, data)
	_, err := os.Stdout.Write([]byte(str))
	return err
}

// newFileWriter 创建文件写入器实例
func newFileWriter() *fileWriter {
	return &fileWriter{}
}

// 创建命令行写入器实例
func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}
