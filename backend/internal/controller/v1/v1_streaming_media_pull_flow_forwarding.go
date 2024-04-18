package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/streaming_media"
)

func (c *ControllerStreaming_media) PullFlowForwarding(ctx context.Context, req *streaming_media.PullFlowForwardingReq) (res *streaming_media.PullFlowForwardingRes, err error) {
	res = &streaming_media.PullFlowForwardingRes{}
	forwarding, err := service.StreamingMedia().PullFlowForwarding(ctx, req)
	liberr.ErrIsNil(ctx, err, "注册拉流转发失败")
	res = forwarding
	service.UserCtx().GetCtx(ctx).Message = "注册拉流转发成功"
	return
}
