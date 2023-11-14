package initialize

import (
	"context"
	"fmt"
	"gorm.io/gorm"

	systemModel "study_gva/model/system"
	"study_gva/service/system"
)

const initOrderEnsureTables = system.InitOrderExternal - 1

type ensureTables struct{}

// auto run
func init() {
	fmt.Println("------ ensure_tables ----")
	system.RegisterInit(initOrderEnsureTables, &ensureTables{})
}

func (e ensureTables) InitializerName() string {
	return "ensure_tables_created"
}

func (e ensureTables) MigrateTable(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	tables := []interface{}{
		systemModel.SysApi{},
	}
	for _, t := range tables {
		_ = db.AutoMigrate(&t)
		// 视图 authority_menu 会被当成表来创建，引发冲突错误（更新版本的gorm似乎不会）
		// 由于 AutoMigrate() 基本无需考虑错误，因此显式忽略
	}
	return ctx, nil
}

func (e ensureTables) InitializeData(ctx context.Context) (next context.Context, err error) {
	return ctx, nil
}

func (e ensureTables) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	tables := []interface{}{
		systemModel.SysApi{},
	}
	yes := true
	for _, t := range tables {
		yes = yes && db.Migrator().HasTable(t)
	}
	return yes
}

func (e ensureTables) DataInserted(ctx context.Context) bool {
	return true
}
