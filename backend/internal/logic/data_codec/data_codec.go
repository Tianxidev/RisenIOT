package dataCodec

import "backend/internal/service"

type sDataCodec struct {
}

func init() {
	service.RegisterDataCodec(New())
}

func New() service.IDataCodec {
	return &sDataCodec{}
}
