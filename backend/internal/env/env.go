package env

import (
	"gopkg.in/ini.v1"
)

// GetEnv 获取环境变量
func GetEnv(key string) (string, error) {

	cfg, err := ini.Load(".env")
	if err != nil {
		return "", err
	}
	service, err := cfg.GetSection("")
	if err != nil {
		return "", err
	}
	startKey, err := service.GetKey(key)
	if err != nil {
		return "", err
	}
	return startKey.Value(), nil

}

// SetEnv 设置环境变量
func SetEnv(key string, value string) error {
	cfg, err := ini.Load(".env")
	if err != nil {
		return err
	}
	service, err := cfg.GetSection("")
	if err != nil {
		return err
	}
	targetKey := service.Key(key)
	if err != nil {
		return err
	}
	targetKey.SetValue(value)
	return cfg.SaveTo(".env")
}
