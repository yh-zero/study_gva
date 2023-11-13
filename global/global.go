package global

import (
	"go.uber.org/zap"
	"study_gva/config"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
)

var (
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger

	BlackCache local_cache.Cache
)
