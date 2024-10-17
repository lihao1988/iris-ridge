package indoor

import (
	"ridge/common/global"
	"ridge/common/lib/db"
)

// AutoMigrate auto exec migration
func AutoMigrate() error {
	// migrate table model(struct by gorm)
	err := db.MigTableModel(global.GDB, global.GModels...)
	if err != nil {
		return err
	}
	global.GModels = nil // 释放变量

	// auto migration for goose
	gConfig := global.GConfig
	err = db.MigGoose(gConfig)
	if err != nil {
		return err
	}

	return nil
}
