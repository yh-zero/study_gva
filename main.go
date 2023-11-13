package main

import (
	"go.uber.org/zap"
	"study_gva/core"
	"study_gva/global"
	"study_gva/initialize"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)

}
