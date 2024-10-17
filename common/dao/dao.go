package dao

import (
	"gorm.io/gorm"
	"ridge/common/lib/db"
	"ridge/common/model"
)

// Dao the base dao
type Dao struct {
	// gorm_db instance
	db *gorm.DB
}

// Create new the data record
func (d *Dao) Create(m model.Impl) error {
	return db.Create(m)
}

// Update modify the data record
func (d *Dao) Update(m model.Impl) error {
	return db.Update(m)
}

// First get the first record
func (d *Dao) First(m model.Impl, options ...func(*gorm.DB) *gorm.DB) error {
	return db.First(m, options...)
}

// Find get the record
func (d *Dao) Find(m model.Impl, options ...func(*gorm.DB) *gorm.DB) error {
	return db.Find(m, options...)
}

// List get the list of record
func (d *Dao) List(b model.Impl, res interface{}, options ...func(*gorm.DB) *gorm.DB) error {
	return db.List(b, res, options...)
}

// PageList get the list of record by page
func (d *Dao) PageList(m model.Impl, res interface{}, page, pageSize int, options ...func(*gorm.DB) *gorm.DB) (int64, error) {
	return db.PageList(m, res, page, pageSize, options...)
}

// Delete delete the record
func (d *Dao) Delete(m model.Impl, options ...func(*gorm.DB) *gorm.DB) error {
	return db.Delete(m, options...)
}
