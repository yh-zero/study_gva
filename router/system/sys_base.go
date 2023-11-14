package system

import (
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	baseRouter := Router.Group("base")
	//baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	//{
	//
	//}
	return baseRouter
}
