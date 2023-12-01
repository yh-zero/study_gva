package response

import "study_gva/model/system"

type SysAPIResponse struct {
	Api system.SysApi `json:"api"`
}

type SysAPIListRespinse struct {
	Apis []system.SysApi `json:"apis"`
}
