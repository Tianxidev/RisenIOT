package sysUserOnline

import (
	"backend/api/v1/common"
	"backend/api/v1/system"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/model"
	"backend/internal/model/do"
	"backend/internal/model/entity"
	"backend/internal/service"
	"backend/library/liberr"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mssola/user_agent"
)

type sSysUserOnline struct {
	Pool *grpool.Pool
}

func init() {
	service.RegisterSysUserOnline(New())
}

func New() *sSysUserOnline {
	return &sSysUserOnline{
		Pool: grpool.New(100),
	}
}

// Invoke 用户在线状态
func (s *sSysUserOnline) Invoke(ctx context.Context, params *model.SysUserOnlineParams) {
	s.Pool.Add(ctx, func(ctx context.Context) {
		//写入数据
		s.SaveOnline(ctx, params)
	})
}

// SaveOnline 保存用户在线状态
func (s *sSysUserOnline) SaveOnline(ctx context.Context, params *model.SysUserOnlineParams) {
	err := g.Try(ctx, func(ctx context.Context) {

		// 获取当前用户id
		uid := service.UserCtx().GetUserId(ctx)

		ua := user_agent.New(params.UserAgent)
		//ua := ""
		browser, _ := ua.Browser()
		//browser, _ := ua.Browser()
		os := ua.OS()
		var (
			info *entity.SysUserOnline
			data = &do.SysUserOnline{
				Uuid:       params.Uuid,
				Token:      params.Token,
				CreateTime: gtime.Now(),
				UserName:   params.Username,
				UserId:     int(uid),
				Ip:         params.Ip,
				Explorer:   browser, // 浏览器
				Os:         os,      // 操作系统
			}
		)

		//查询是否已存在当前用户
		err := dao.SysUserOnline.Ctx(ctx).Fields(dao.SysUserOnline.Columns().Id).
			Where(dao.SysUserOnline.Columns().Token, data.Token).
			Scan(&info)
		liberr.ErrIsNil(ctx, err)
		//若已存在则更新
		if info != nil {
			_, err = dao.SysUserOnline.Ctx(ctx).
				Where(dao.SysUserOnline.Columns().Id, info.Id).
				FieldsEx(dao.SysUserOnline.Columns().Id).Update(data)
			liberr.ErrIsNil(ctx, err)
		} else { //否则新增
			_, err = dao.SysUserOnline.Ctx(ctx).
				FieldsEx(dao.SysUserOnline.Columns().Id).Insert(data)
			liberr.ErrIsNil(ctx, err)
		}
	})
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

// CheckUserOnline 检查在线用户
func (s *sSysUserOnline) CheckUserOnline(ctx context.Context) {
	param := &system.SysUserOnlineSearchReq{
		PageReq: common.PageReq{
			PageNum:  1,
			PageSize: 50,
		},
	}
	var total int
	for {
		var (
			res *system.SysUserOnlineSearchRes
			err error
		)
		res, err = s.GetOnlineListPage(ctx, param, true)
		if err != nil {
			g.Log().Error(ctx, err)
			break
		}
		if res.List == nil {
			break
		}
		for _, v := range res.List {
			if b := s.UserIsOnline(ctx, v.Token); !b {
				s.DeleteOnlineByToken(ctx, v.Token)
			}
		}
		if param.PageNum*param.PageSize >= total {
			break
		}
		param.PageNum++
	}
}

// GetOnlineListPage 搜素在线用户列表
func (s *sSysUserOnline) GetOnlineListPage(ctx context.Context, req *system.SysUserOnlineSearchReq, hasToken ...bool) (res *system.SysUserOnlineSearchRes, err error) {
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	model := dao.SysUserOnline.Ctx(ctx)
	if req.Ip != "" {
		model = model.Where("ip like ?", "%"+req.Ip+"%")
	}
	if req.Username != "" {
		model = model.Where("user_name like ?", "%"+req.Username+"%")
	}
	res = new(system.SysUserOnlineSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = model.Count()
		liberr.ErrIsNil(ctx, err, "获取总行数失败")
		if len(hasToken) == 0 || !hasToken[0] {
			model = model.FieldsEx("token")
		}
		err = model.Page(req.PageNum, req.PageSize).Order("create_time DESC").Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

func (s *sSysUserOnline) UserIsOnline(ctx context.Context, token string) bool {
	err := g.Try(ctx, func(ctx context.Context) {
		_, _, err := service.Token().Get().GetTokenData(ctx, token)
		liberr.ErrIsNil(ctx, err)
	})
	return err == nil
}

func (s *sSysUserOnline) DeleteOnlineByToken(ctx context.Context, token string) (err error) {
	_, err = dao.SysUserOnline.Ctx(ctx).Delete(dao.SysUserOnline.Columns().Token, token)
	return
}

// ForceLogout 强制用户下线
func (s *sSysUserOnline) ForceLogout(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var onlineList []*entity.SysUserOnline
		onlineList, err = s.GetInfosByIds(ctx, ids)
		liberr.ErrIsNil(ctx, err)
		_, err = dao.SysUserOnline.Ctx(ctx).Where(dao.SysUserOnline.Columns().Id+" in(?)", ids).Delete()
		liberr.ErrIsNil(ctx, err)
		for _, v := range onlineList {
			err = service.Token().Get().RemoveToken(ctx, v.Token)
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

// ForceLogoutAll 强制所有用户下线
func (s *sSysUserOnline) ForceLogoutAll(ctx context.Context) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 查询所有在线用户
		var onlineList []*entity.SysUserOnline
		err = dao.SysUserOnline.Ctx(ctx).Scan(&onlineList)
		liberr.ErrIsNil(ctx, err)
		for _, v := range onlineList {
			_ = service.Token().Get().RemoveToken(ctx, v.Token)
		}

		_, err = dao.SysUserOnline.Ctx(ctx).Wheref(`? != ?`, dao.SysUserOnline.Columns().UserId, service.UserCtx().GetUserId(ctx)).Delete()
		liberr.ErrIsNil(ctx, err, "强制所有用户下线失败")
	})
	return
}

// GetInfosByIds 根据ids获取用户在线信息
func (s *sSysUserOnline) GetInfosByIds(ctx context.Context, ids []int) (onlineList []*entity.SysUserOnline, err error) {
	err = dao.SysUserOnline.Ctx(ctx).Where(dao.SysUserOnline.Columns().Id+" in(?)", ids).Scan(&onlineList)
	return
}
