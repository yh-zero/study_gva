package router

import (
	"study_gva/router/example"
	"study_gva/router/system"
	"study_gva/router/test001"
	"study_gva/router/test002"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	TEST001 test001.RouterGroup
	Test002 test002.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
