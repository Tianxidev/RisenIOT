package env

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

// GetEnv 获取环境变量
func GetEnv(key string) (string, error) {
	cfg, err := ini.Load(".env")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	return cfg.Section("").Key(key).String(), nil
}
