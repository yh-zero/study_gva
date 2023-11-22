package system

import (
	"go.uber.org/zap"

	"study_gva/global"
	"study_gva/model/common/response"
	"study_gva/model/system"
	systemReq "study_gva/model/system/request"

	"github.com/gin-gonic/gin"
)

type OperationRecordApi struct{}

// 删除操作记录
func (s *OperationRecordApi) DeleteSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	err := c.ShouldBindJSON(&sysOperationRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationRecordService.DeleteSysOperationRecord(sysOperationRecord)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// 分页获取SysOperationRecord列表 操作记录列表
func (s *OperationRecordApi) GetSysOperationRecordList(c *gin.Context) {
	var pageInfo systemReq.SysOperationRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := operationRecordService.GetSysOperationRecordInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)

}
