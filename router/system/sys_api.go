package system

import (
	"github.com/gin-gonic/gin"
	v1 "study_gva/api/v1"
	"study_gva/middleware"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("api")

	apiPublicRouterWithoutRecord := RouterPub.Group("api")
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		apiRouter.POST("createApi", apiRouterApi.CreateApi)               // 创建api
		apiRouter.POST("getApiById", apiRouterApi.GetApiById)             // 获取单条Api消息
		apiRouter.POST("updateApi", apiRouterApi.UpdateApi)               // 更新api
		apiRouter.POST("deleteApi", apiRouterApi.DeleteApi)               // 删除api
		apiRouter.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds) // ids 同时删除多条数据
	}

	{
		apiRouterWithoutRecord.POST("getApiList", apiRouterApi.GetApiList) // 获取Api列表
		apiRouterWithoutRecord.POST("getAllApis", apiRouterApi.GetAllApis) // 获取api列表
	}

	{
		apiPublicRouterWithoutRecord.GET("freshCasbin", apiRouterApi.FreshCasbin) // 刷新casbin权限
	}
}
