// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/internal/model"
	"context"
)

type (
	IDataCodec interface {
		// HttpDecode http解码
		HttpDecode(ctx context.Context, dataContent interface{}) (dmesg *model.DeviceDecodeMsg, err error)
		// Save 保存数据
		Save(ctx context.Context, msg *model.DeviceDecodeMsg) (err error)
	}
)

var (
	localDataCodec IDataCodec
)

func DataCodec() IDataCodec {
	if localDataCodec == nil {
		panic("implement not found for interface IDataCodec, forgot register?")
	}
	return localDataCodec
}

func RegisterDataCodec(i IDataCodec) {
	localDataCodec = i
}
