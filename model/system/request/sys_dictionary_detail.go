package request

import (
	"study_gva/model/common/request"
	"study_gva/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
