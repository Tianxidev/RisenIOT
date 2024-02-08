// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/api/v1/system"
	"backend/internal/model"
	"context"
)

type (
	IPersonal interface {
		GetPersonalInfo(ctx context.Context, req *system.PersonalInfoReq) (res *system.PersonalInfoRes, err error)
		EditPersonal(ctx context.Context, req *system.PersonalEditReq) (user *model.LoginUserRes, err error)
		ResetPwdPersonal(ctx context.Context, req *system.PersonalResetPwdReq) (res *system.PersonalResetPwdRes, err error)
	}
)

var (
	localPersonal IPersonal
)

func Personal() IPersonal {
	if localPersonal == nil {
		panic("implement not found for interface IPersonal, forgot register?")
	}
	return localPersonal
}

func RegisterPersonal(i IPersonal) {
	localPersonal = i
}
