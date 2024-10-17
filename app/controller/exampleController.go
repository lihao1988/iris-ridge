package controller

import (
	"ridge/common/global"
)

// ExampleController controller demo
type ExampleController struct {
	Controller
}

// init append "Example" to GControllers
func init() {
	// add register controller
	global.AddControllerRegister(new(ExampleController))
}

// Prefix custom routing prefix of "Example" controller
func (p *ExampleController) Prefix() string {
	return "example"
}

// BuildRoute build the controller route
func (p *ExampleController) BuildRoute() {
	p.Mvc.Handle(p)
}

// GetHw demo function
func (p *ExampleController) GetHw() {
	p.Msg(200, "hello world !", "success")
}

// GetDemo for view
func (p *ExampleController) GetDemo() {
	_ = p.Ctx.View("demo.html")
}
