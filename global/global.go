package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	"study_gva/config"
	"study_gva/utils/timer"

	"github.com/qiniu/qmgo"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/songzhibin97/gkit/cache/singleflight"
	"github.com/spf13/viper"
)

var (
	GVA_CONFIG              config.Server
	GVA_VP                  *viper.Viper
	GVA_LOG                 *zap.Logger
	GVA_DB                  *gorm.DB
	GVA_DBList              map[string]*gorm.DB
	GVA_Timer               timer.Timer = timer.NewTimerTask()
	GVA_REDIS               *redis.Client
	GVA_MONGO               *qmgo.QmgoClient
	GVA_Concurrency_Control = &singleflight.Group{}

	BlackCache local_cache.Cache
)
