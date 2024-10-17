package model

import "gorm.io/gorm"

// Impl the model interface
// use gorm
type Impl interface {
	// TableName TableComment migration function
	TableName() string
	TableComment() string

	// SetDB GetDB SetID GetID business function
	SetDB(db *gorm.DB)
	GetDB() *gorm.DB
	SetID(id uint)
	GetID() uint
}
