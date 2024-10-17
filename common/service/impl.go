package service

import (
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

// Impl the service interface
type Impl interface {
	// SetCtx GetCtx iris context
	SetCtx(ctx iris.Context)
	GetCtx() iris.Context

	// SetDB GetDB gorm_db
	SetDB(db *gorm.DB)
	GetDB() *gorm.DB
}
