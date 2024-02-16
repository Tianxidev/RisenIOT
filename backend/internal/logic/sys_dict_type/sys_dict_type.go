package sysDictType

import (
	"backend/api/v1/system"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/model"
	"backend/internal/model/do"
	"backend/internal/model/entity"
	"backend/internal/service"
	"backend/library/liberr"
	"context"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysDictType struct {
}

func init() {
	service.RegisterSysDictType(New())
}

func New() *sSysDictType {
	return &sSysDictType{}
}

// List 字典类型列表
func (s *sSysDictType) List(ctx context.Context, req *system.DictTypeSearchReq) (res *system.DictTypeSearchRes, err error) {
	res = new(system.DictTypeSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysDictType.Ctx(ctx)
		if req.DictName != "" {
			m = m.Where(dao.SysDictType.Columns().DictName+" like ?", "%"+req.DictName+"%")
		}
		if req.DictType != "" {
			m = m.Where(dao.SysDictType.Columns().DictType+" like ?", "%"+req.DictType+"%")
		}
		if req.Status != "" {
			m = m.Where(dao.SysDictType.Columns().Status+" = ", gconv.Int(req.Status))
		}
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取字典类型失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		res.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		err = m.Fields(model.SysDictTypeInfoRes{}).Page(req.PageNum, req.PageSize).
			Order(dao.SysDictType.Columns().DictId + " asc").Scan(&res.DictTypeList)
		liberr.ErrIsNil(ctx, err, "获取字典类型失败")
	})
	return
}

// Add 添加字典类型
func (s *sSysDictType) Add(ctx context.Context, req *system.DictTypeAddReq, userId uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = s.ExistsDictType(ctx, req.DictType)
		liberr.ErrIsNil(ctx, err)
		_, err = dao.SysDictType.Ctx(ctx).Insert(do.SysDictType{
			DictName: req.DictName,
			DictType: req.DictType,
			Status:   req.Status,
			CreateBy: userId,
			Remark:   req.Remark,
		})
		liberr.ErrIsNil(ctx, err, "添加字典类型失败")
		//清除缓存
		service.Cache().GCache().RemoveByTag(ctx, consts.CacheSysDictTag)
	})
	return
}

// Edit 修改字典类型
func (s *sSysDictType) Edit(ctx context.Context, req *system.DictTypeEditReq, userId uint64) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			err = s.ExistsDictType(ctx, req.DictType, req.DictId)
			liberr.ErrIsNil(ctx, err)
			dictType := (*entity.SysDictType)(nil)
			e := dao.SysDictType.Ctx(ctx).Fields(dao.SysDictType.Columns().DictType).WherePri(req.DictId).Scan(&dictType)
			liberr.ErrIsNil(ctx, e, "获取字典类型失败")
			liberr.ValueIsNil(ctx, dictType, "字典类型不存在")
			// 修改字典类型
			_, e = dao.SysDictType.Ctx(ctx).TX(tx).WherePri(req.DictId).Update(do.SysDictType{
				DictName: req.DictName,
				DictType: req.DictType,
				Status:   req.Status,
				UpdateBy: userId,
				Remark:   req.Remark,
			})
			liberr.ErrIsNil(ctx, e, "修改字典类型失败")
			// 修改字典数据
			_, e = dao.SysDictData.Ctx(ctx).TX(tx).Data(do.SysDictData{DictType: req.DictType}).
				Where(dao.SysDictData.Columns().DictType, dictType.DictType).Update()
			liberr.ErrIsNil(ctx, e, "修改字典数据失败")
			// 清除缓存
			service.Cache().GCache().RemoveByTag(ctx, consts.CacheSysDictTag)
		})
		return err
	})
	return
}

func (s *sSysDictType) Get(ctx context.Context, req *system.DictTypeGetReq) (dictType *entity.SysDictType, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysDictType.Ctx(ctx).Where(dao.SysDictType.Columns().DictId, req.DictId).Scan(&dictType)
		liberr.ErrIsNil(ctx, err, "获取字典类型失败")
	})
	return
}

// ExistsDictType 检查类型是否已经存在
func (s *sSysDictType) ExistsDictType(ctx context.Context, dictType string, dictId ...int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysDictType.Ctx(ctx).Fields(dao.SysDictType.Columns().DictId).
			Where(dao.SysDictType.Columns().DictType, dictType)
		if len(dictId) > 0 {
			m = m.Where(dao.SysDictType.Columns().DictId+" !=? ", dictId[0])
		}
		res, e := m.One()
		liberr.ErrIsNil(ctx, e, "sql err")
		if !res.IsEmpty() {
			liberr.ErrIsNil(ctx, gerror.New("字典类型已存在"))
		}
	})
	return
}

// Delete 删除字典类型
func (s *sSysDictType) Delete(ctx context.Context, dictIds []int) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			discs := ([]*entity.SysDictType)(nil)
			err = dao.SysDictType.Ctx(ctx).Fields(dao.SysDictType.Columns().DictType).
				Where(dao.SysDictType.Columns().DictId+" in (?) ", dictIds).Scan(&discs)
			liberr.ErrIsNil(ctx, err, "删除失败")
			types := garray.NewStrArray()
			for _, dt := range discs {
				types.Append(dt.DictType)
			}
			if types.Len() > 0 {
				_, err = dao.SysDictType.Ctx(ctx).TX(tx).Delete(dao.SysDictType.Columns().DictId+" in (?) ", dictIds)
				liberr.ErrIsNil(ctx, err, "删除类型失败")
				_, err = dao.SysDictData.Ctx(ctx).TX(tx).Delete(dao.SysDictData.Columns().DictType+" in (?) ", types.Slice())
				liberr.ErrIsNil(ctx, err, "删除字典数据失败")
			}
			//清除缓存
			service.Cache().GCache().RemoveByTag(ctx, consts.CacheSysDictTag)
		})
		return err
	})
	return
}

// GetAllDictType 获取所有正常状态下的字典类型
func (s *sSysDictType) GetAllDictType(ctx context.Context) (list []*entity.SysDictType, err error) {
	cache := service.Cache()
	//从缓存获取
	data := cache.GCache().Get(ctx, consts.CacheSysDict+"_dict_type_all")
	if !data.IsNil() {
		err = data.Structs(&list)
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysDictType.Ctx(ctx).Where("status", 1).Order("dict_id ASC").Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取字典类型数据出错")
		//缓存
		cache.GCache().Set(ctx, consts.CacheSysDict+"_dict_type_all", list, 0, consts.CacheSysDictTag)
	})
	return
}
