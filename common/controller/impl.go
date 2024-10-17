package controller

import (
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

// Impl the controller interface
type Impl interface {
	// Party Prefix route data function
	Party() string  // custom routing prefix of module
	Prefix() string // custom routing prefix of controller

	// GetDB get db of gorm
	GetDB() *gorm.DB

	// BuildMvc BuildRoute Msg Send business function
	BuildMvc(app *iris.Application, controllerName string) // build controller mvc
	BuildRoute()                                           // build controller route
	Msg(code int, data interface{}, msg string)            // output in json
	Send(data interface{}, err error, msg ...string)       // output of data
}
