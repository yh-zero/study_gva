package service

import (
	"study_gva/service/system"
	"study_gva/service/test001"
	"study_gva/service/test002"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	TEST001ServiceGroup test001.ServiceGroup
	Test002ServiceGroup test002.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
