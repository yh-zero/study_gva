package response

import "study_gva/model/system"

type SysMenusResponse struct {
	Menus []system.SysMenu `json:"menus"`
}