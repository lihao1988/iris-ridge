package global

import (
	"ridge/common/controller"
	"ridge/common/model"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var (
	// GApplication iris App
	GApplication *iris.Application

	// GValidator request validate
	GValidator   *validator.Validate
	GVTranslator ut.Translator

	// SessionManager session auth
	SessionManager *sessions.Sessions

	// GControllers all controller set
	GControllers []controller.Impl

	// GModels all table model set
	GModels []model.Impl
)

// AddControllerRegister add controller for register router
func AddControllerRegister(c controller.Impl) {
	GControllers = append(GControllers, c)
}

// AddModelRegister add model for register
func AddModelRegister(m model.Impl) {
	GModels = append(GModels, m)
}
