package logger

import (
	"fmt"
	"time"
)

// GetTimeString 获取时间字符串
func GetTimeString() string {

	Y := time.Now().Format("2006")
	M := time.Now().Format("01")
	D := time.Now().Format("02")
	h := time.Now().Format("15")
	m := time.Now().Format("04")
	s := time.Now().Format("05")

	return fmt.Sprintf("%s-%s-%s %s:%s:%s", Y, M, D, h, m, s)
}
