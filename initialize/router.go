package initialize

import (
	"github.com/gin-gonic/gin"
	"study_gva/global"
)

// 初始化总路由

func Routers() *gin.Engine {
	// 设置为发布模式
	if global.GVA_CONFIG.System.Env == "public" {
		gin.SetMode(gin.ReleaseMode)
	}

	Router := gin.New()

	if global.GVA_CONFIG.System.Env != "public" {
		Router.Use(gin.Logger(), gin.Recovery())
	}

	InstallPlugin(Router) // 安装插件

}
