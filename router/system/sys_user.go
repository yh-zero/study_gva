package system

import (
	"github.com/gin-gonic/gin"
	v1 "study_gva/api/v1"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		//无需操作记录
	}

	{
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo) // 获取自身信息
	}

}
