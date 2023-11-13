package request

import (
	"study_gva/model/common/request"
	"study_gva/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
