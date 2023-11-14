package system

import "study_gva/service"

type ApiGroup struct {
	BaseApi
	DBApi
}

var (
	initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
)
