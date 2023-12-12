package system

import (
	"github.com/gin-gonic/gin"
	v1 "study_gva/api/v1"
	"study_gva/middleware"
)

type CasbinRouter struct{}

func (s *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) {
	casbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	casbinRouterWithoutRecord := Router.Group("casbin")
	casbinApi := v1.ApiGroupApp.SystemApiGroup.CasbinApi

	{
		casbinRouter.POST("updateCasbin", casbinApi.UpdateCasbin)
	}

	{
		casbinRouterWithoutRecord.POST("getPolicyPathByAuthorityId", casbinApi.GetPolicyPathByAuthorityId)
	}
}
