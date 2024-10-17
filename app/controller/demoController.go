package controller

import (
	"ridge/common/global"
)

// DemoController controller demo
type DemoController struct {
	Controller
}

// init append "Demo" to GControllers
func init() {
	// add register controller
	global.AddControllerRegister(new(DemoController))
}

// Prefix custom routing prefix of "Demo" controller
func (d *DemoController) Prefix() string {
	return "demo"
}

// BuildRoute build the controller route
func (d *DemoController) BuildRoute() {
	d.Mvc.Handle(d)
}

// GetHw demo function
func (d *DemoController) GetHw() {
	d.Msg(200, "demo, hello world !", "success")
}

// GetTest for view test
func (d *DemoController) GetTest() {
	_ = d.Ctx.View("test/test.html")
}
