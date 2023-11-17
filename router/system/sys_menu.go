package system

import (
	"github.com/gin-gonic/gin"
	v1 "study_gva/api/v1"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	//menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("menu")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.AuthorityMenuApi

	{
		menuRouterWithoutRecord.POST("getMenu", authorityMenuApi.GetMenu) // 获取菜单树
	}
	return menuRouterWithoutRecord
}
