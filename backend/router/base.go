package router

import (
	"RisenIOT/backend/controller"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	r := Router.Group("base")
	r.GET("/version", controller.GetVersion)
}
