// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/api/v1/streaming_media"
	"context"
)

type (
	IStreamingMedia interface {
		// PullFlowForwarding 拉流转发服务
		PullFlowForwarding(ctx context.Context, req *streaming_media.PullFlowForwardingReq) (res *streaming_media.PullFlowForwardingRes, err error)
	}
)

var (
	localStreamingMedia IStreamingMedia
)

func StreamingMedia() IStreamingMedia {
	if localStreamingMedia == nil {
		panic("implement not found for interface IStreamingMedia, forgot register?")
	}
	return localStreamingMedia
}

func RegisterStreamingMedia(i IStreamingMedia) {
	localStreamingMedia = i
}
