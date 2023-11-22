package system

import (
	v1 "study_gva/api/v1"

	"github.com/gin-gonic/gin"
)

type OperationRecordRouter struct{}

func (s *OperationRecordRouter) InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("sysOperationRecord")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.OperationRecordApi

	{
		operationRecordRouter.GET("getSysOperationRecordList", authorityMenuApi.GetSysOperationRecordList)  // 获取SysOperationRecord列表
		operationRecordRouter.DELETE("deleteSysOperationRecord", authorityMenuApi.DeleteSysOperationRecord) // 删除SysOperationRecord
	}

}
