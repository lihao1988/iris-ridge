package db

import (
	"errors"
	"fmt"

	baseModel "ridge/common/model"

	"gorm.io/gorm"
)

// MigTableModel migrate table model struct
func MigTableModel(db *gorm.DB, tbModel ...baseModel.Impl) error {
	if db == nil {
		return errors.New("DB is nil ~")
	}

	// inti table struct
	var err error
	for _, tb := range tbModel {
		// table for create or update
		optionData := fmt.Sprintf("ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 "+
			"COLLATE=utf8mb4_general_ci comment '%s'", tb.TableComment())
		err = db.Set("gorm:table_options", optionData).Migrator().AutoMigrate(tb)
		if err != nil {
			return err
		}
	}

	return nil
}
