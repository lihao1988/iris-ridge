package dao

import (
	"ridge/common/model"

	"gorm.io/gorm"
)

// Impl the dao interface
type Impl interface {
	// Create new the data record
	Create(m model.Impl) error

	// Update modify the data record
	Update(m model.Impl) error

	// First get the first record
	First(m model.Impl, options ...func(*gorm.DB) *gorm.DB) error
	// Find get the record
	Find(m model.Impl, options ...func(*gorm.DB) *gorm.DB) error
	// List get the list of record
	List(b model.Impl, res interface{}, options ...func(*gorm.DB) *gorm.DB) error
	// PageList get the list of record by page
	PageList(m model.Impl, res interface{}, page, pageSize int, options ...func(*gorm.DB) *gorm.DB) (int64, error)

	// Delete delete the record
	Delete(m model.Impl, options ...func(*gorm.DB) *gorm.DB) error
}
