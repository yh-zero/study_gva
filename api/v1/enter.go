package v1

import (
	"study_gva/api/v1/system"
	"study_gva/api/v1/test001"
	"study_gva/api/v1/test002"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	TEST001ApiGroup test001.ApiGroup
	Test002ApiGroup test002.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
