package service

import (
	"ridge/common/global"

	"gorm.io/gorm"

	"github.com/kataras/iris/v12"
)

// Service the base service
type Service struct {
	// iris context
	ctx iris.Context

	// gorm_db instance
	db *gorm.DB
}

// SetCtx set iris context
func (d *Service) SetCtx(ctx iris.Context) {
	d.ctx = ctx
}

// GetCtx get iris context
func (d *Service) GetCtx() iris.Context {
	return d.ctx
}

// SetDB set db
func (d *Service) SetDB(db *gorm.DB) {
	d.db = db
}

// GetDB get db
func (d *Service) GetDB() *gorm.DB {
	if d.db != nil {
		return d.db
	}

	return global.GDB
}
