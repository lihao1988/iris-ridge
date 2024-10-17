package controller

import (
	"ridge/common/lib/response"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"gorm.io/gorm"
)

// Controller the base controller struct
type Controller struct {
	Ctx iris.Context
	Mvc *mvc.Application
}

// Party custom routing prefix of module
func (c *Controller) Party() string {
	return ""
}

// Prefix custom routing prefix of controller
func (c *Controller) Prefix() string {
	return ""
}

// GetDB get db of gorm
func (c *Controller) GetDB() *gorm.DB {
	return nil
}

// BuildMvc build the controller MVC
func (c *Controller) BuildMvc(app *iris.Application, controllerName string) {
	controllerParty := app.Party("/" + controllerName)
	c.Mvc = mvc.New(controllerParty)
	if c.Ctx != nil {
		c.Mvc.Register(c.Ctx)
	}
}

// BuildRoute build the controller route
func (c *Controller) BuildRoute() {
	c.Mvc.Handle(c)
}

// Msg output in json
func (c *Controller) Msg(code int, data interface{}, msg string) {
	// output message
	response.Msg(c.Ctx, code, data, msg)
	return
}

// Send output of data
func (c *Controller) Send(data interface{}, err error, msg ...string) {
	// send message
	response.SendMsg(c.Ctx, data, err, msg...)
	return
}
