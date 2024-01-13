// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package login

import (
	"context"

	"backend/api/login/v1"
)

type ILoginV1 interface {
	UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error)
}
