package indoor

import (
	"fmt"
	rApp "ridge/app"
	"ridge/common/global"
	"ridge/route"

	"github.com/kataras/iris/v12"
)

// AutoloadRoute autoload route set
func AutoloadRoute() error {
	// iris app
	app := global.GApplication

	// the direct route
	BuildRoute(app)

	// build controller route
	// whether auto init Mvc route
	gConfig := global.GConfig
	if gConfig.App.WithRoute {
		err := MvcRoute(app)
		if err != nil {
			return err
		}
	}

	// custom controller route
	rApp.BuildRoute(app)

	return nil
}

// BuildRoute build the direct route
func BuildRoute(app *iris.Application) {
	route.Build(app)
	return
}

// MvcRoute build the route for controller
func MvcRoute(app *iris.Application) error {
	// build mvc roure
	var routePath, moduleName string
	for _, cont := range global.GControllers {
		// route prefix
		routePath = cont.Prefix()
		moduleName = cont.Party()
		if moduleName != "" {
			routePath = fmt.Sprintf("%s/%s", moduleName, routePath)
		}

		// the "BuildMvc" function of controller
		cont.BuildMvc(app, routePath)

		// build route
		cont.BuildRoute()
	}
	return nil
}
