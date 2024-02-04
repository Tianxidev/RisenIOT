// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/api/v1/system"
	"backend/internal/model/entity"
	"context"
)

type (
	ISysPost interface {
		// List 岗位列表
		List(ctx context.Context, req *system.PostSearchReq) (res *system.PostSearchRes, err error)
		Add(ctx context.Context, req *system.PostAddReq) (err error)
		Edit(ctx context.Context, req *system.PostEditReq) (err error)
		Delete(ctx context.Context, ids []int) (err error)
		// GetUsedPost 获取正常状态的岗位
		GetUsedPost(ctx context.Context) (list []*entity.SysPost, err error)
	}
)

var (
	localSysPost ISysPost
)

func SysPost() ISysPost {
	if localSysPost == nil {
		panic("implement not found for interface ISysPost, forgot register?")
	}
	return localSysPost
}

func RegisterSysPost(i ISysPost) {
	localSysPost = i
}
