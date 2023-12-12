package system

import (
	"github.com/gin-gonic/gin"
	v1 "study_gva/api/v1"
)

type AutoCodeRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeRouter(Router *gin.RouterGroup) {
	autoCodeRouter := Router.Group("autoCode")
	autoCodeApi := v1.ApiGroupApp.SystemApiGroup.AutoCodeApi
	{
		autoCodeRouter.POST("getPackage", autoCodeApi.GetPackage)       // 获取package包
		autoCodeRouter.POST("createPackage", autoCodeApi.CreatePackage) // 创建package包
		autoCodeRouter.POST("delPackage", autoCodeApi.DelPackage)       // 删除package包
	}

}
