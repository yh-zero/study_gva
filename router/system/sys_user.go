package system

import (
	"github.com/gin-gonic/gin"
	v1 "study_gva/api/v1"
	"study_gva/middleware"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("admin_register", baseApi.Register)     // 管理员注册账号
		userRouter.POST("resetPassword", baseApi.ResetPassword) // 重置密码
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)      // 设置用户信息
		userRouter.DELETE("deleteUser", baseApi.DeleteUser)     // 删除用户
	}

	{
		//无需操作记录
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)  // 获取自身信息
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList) // 分页获取用户列表
	}

}
