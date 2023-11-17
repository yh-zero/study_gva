package system

import (
	"go.uber.org/zap"

	"study_gva/global"
	"study_gva/model/common/response"
	"study_gva/model/system"
	systemRes "study_gva/model/system/response"
	"study_gva/utils"

	"github.com/gin-gonic/gin"
)

type AuthorityMenuApi struct{}

// 获取用户动态路由
func (a *AuthorityMenuApi) GetMenu(c *gin.Context) {
	global.GVA_LOG.Info("GetMenu======")
	menus, err := menuService.GetMenuTree(utils.GetUserAuthorityId(c))

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	if menus == nil {
		menus = []system.SysMenu{}
	}
	response.OkWithDetailed(systemRes.SysMenusResponse{
		Menus: menus,
	}, "获取成功", c)
}
