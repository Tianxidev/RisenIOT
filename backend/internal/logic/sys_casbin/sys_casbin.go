package sysCasbin

import (
	"backend/internal/service"
	"context"
	"github.com/casbin/casbin/v2"
	"sync"
)

type sCasbin struct {
	Enforcer    *casbin.SyncedEnforcer
	EnforcerErr error
	ctx         context.Context
}

var (
	once  sync.Once
	en    *casbin.SyncedEnforcer
	enErr error
)

func init() {
	service.RegisterCasbin(New())
}

func New() *sCasbin {
	return &sCasbin{}
}
