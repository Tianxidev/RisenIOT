package user

import (
	"backend/internal/dao"
	"backend/internal/model/do"
	"backend/internal/model/entity"
	"backend/internal/service"
	"context"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

// GetByUsername 根据用户名获取用户信息
func (s *sUser) GetByUsername(ctx context.Context, username string) (data entity.RiseniotUser, err error) {
	err = dao.RiseniotUser.Ctx(ctx).Where(do.RiseniotUser{
		Username: username,
	}).Scan(&data)
	return data, err
}
