package system

import "study_gva/service"

type ApiGroup struct {
	BaseApi
	DBApi
	AuthorityMenuApi
	AuthorityApi
}

var (
	initDBService    = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	userService      = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService       = service.ServiceGroupApp.SystemServiceGroup.JwtService
	menuService      = service.ServiceGroupApp.SystemServiceGroup.MenuService
	authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	casbinService    = service.ServiceGroupApp.SystemServiceGroup.CasbinService
)
