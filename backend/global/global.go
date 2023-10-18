package global

import (
	"RisenIOT/backend/internal/emqx"
	"RisenIOT/backend/internal/logger"
)

var (
	Logger *logger.Logger
	Emqx   *emqx.Emqx
)
