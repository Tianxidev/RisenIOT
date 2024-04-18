package sysPush

import "backend/internal/service"

type sSysPush struct {
}

func New() *sSysPush {
	return &sSysPush{}
}

func init() {
	service.RegisterSysPush(New())
}
