package controller

import (
	"gorm.io/gorm"
	baseCont "ridge/common/controller"
	"ridge/common/global"
)

// Controller the controller's instance of app
type Controller struct {
	baseCont.Controller
}

// Party custom routing prefix of "app" module
func (c *Controller) Party() string {
	return "app"
}

// GetDB get db of gorm
func (c *Controller) GetDB() *gorm.DB {
	return global.GDB
}
