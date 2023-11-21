package system

import (
	"github.com/gin-gonic/gin"
	v1 "study_gva/api/v1"
	"study_gva/middleware"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("authority").Use(middleware.OperationRecord())
	authorityRouterWithoutRecord := Router.Group("authority")
	authorityApi := v1.ApiGroupApp.SystemApiGroup.AuthorityApi

	{
		authorityRouter.POST("createAuthority", authorityApi.CreateAuthority) // 创建角色
		authorityRouter.PUT("updateAuthority", authorityApi.UpdateAuthority)  // 更新角色
		authorityRouter.POST("deleteAuthority", authorityApi.DeleteAuthority) // 删除角色

	}

	{
		authorityRouterWithoutRecord.POST("getAuthorityList", authorityApi.GetAuthorityList) // 获取角色列表
	}

}
