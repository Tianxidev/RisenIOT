// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ICaptcha interface {
		// GetVerifyImgString 获取字母数字混合验证码
		GetVerifyImgString(ctx context.Context) (idKeyC string, base64stringC string, err error)
		// VerifyString 验证输入的验证码是否正确
		VerifyString(id, answer string) bool
	}
)

var (
	localCaptcha ICaptcha
)

func Captcha() ICaptcha {
	if localCaptcha == nil {
		panic("implement not found for interface ICaptcha, forgot register?")
	}
	return localCaptcha
}

func RegisterCaptcha(i ICaptcha) {
	localCaptcha = i
}
