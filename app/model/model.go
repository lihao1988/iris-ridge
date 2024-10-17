package model

import (
	"gorm.io/gorm"
	"ridge/common/global"
	baseModel "ridge/common/model"
)

// Model the model's instance of app
type Model struct {
	baseModel.Model
}

// GetDB get db
func (m *Model) GetDB() *gorm.DB {
	db := m.Model.GetDB()
	if db != nil {
		return db
	}

	return global.GDB
}
