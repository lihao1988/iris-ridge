package global

import (
	"gorm.io/gorm"
	config "ridge/config/abstract"
)

var (
	// GDB DB
	GDB *gorm.DB

	// GConfig global config
	GConfig = &config.CGlobal{}
)
