package system

import "study_gva/service"

type ApiGroup struct {
	BaseApi
	DBApi
	AuthorityMenuApi
	AuthorityApi
	OperationRecordApi
	DictionaryApi
	DictionaryDetailApi
}

var (
	initDBService           = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	userService             = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService              = service.ServiceGroupApp.SystemServiceGroup.JwtService
	menuService             = service.ServiceGroupApp.SystemServiceGroup.MenuService
	authorityService        = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	casbinService           = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	operationRecordService  = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
	dictionaryService       = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
	dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
)
