package streaming_media

import "backend/internal/service"

type sStreamingMedia struct {
}

func init() {
	service.RegisterStreamingMedia(New())
}

func New() *sStreamingMedia {
	return &sStreamingMedia{}
}
