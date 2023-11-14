package system

import (
	"github.com/gin-gonic/gin"
	v1 "study_gva/api/v1"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("captcha", baseApi.Captcha)
	}
	return baseRouter
}
