package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"study_gva/global"
	"study_gva/model/common/response"
	"study_gva/model/system"
	"study_gva/utils"
)

type AutoCodeApi struct{}

// DelPackage
// @Tags      AutoCode
// @Summary   删除package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAutoCode                                         true  "创建package"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "删除package成功"
// @Router    /autoCode/delPackage [post]
func (autoApi *AutoCodeApi) DelPackage(c *gin.Context) {
	var autoCode system.SysAutoCode
	_ = c.ShouldBindJSON(&autoCode)
	err := autoCodeService.DelPackage(autoCode)
	if err != nil {
		global.GVA_LOG.Error("删除失败！", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// CreatePackage
// @Tags      AutoCode
// @Summary   创建package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAutoCode                                         true  "创建package"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "创建package成功"
// @Router    /autoCode/createPackage [post]
func (autoApi *AutoCodeApi) CreatePackage(c *gin.Context) {
	var autoCode system.SysAutoCode
	_ = c.ShouldBindJSON(&autoCode)

	if err := utils.Verify(autoCode, utils.AutoPackageVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := autoCodeService.CreateAutoCode(&autoCode)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)

}

// GetPackage
// @Tags      AutoCode
// @Summary   获取package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "创建package成功"
// @Router    /autoCode/getPackage [post]
func (autoApi *AutoCodeApi) GetPackage(c *gin.Context) {
	pkgs, err := autoCodeService.GetPackage()
	if err != nil {
		global.GVA_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithDetailed(gin.H{"pkgs": pkgs}, "获取成功", c)

}
