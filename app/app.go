package app

import (
	"fmt"
	"path"
	"strings"

	"ridge/common/global"
	"ridge/tool"

	// for init controller route and table model
	// import all module's controllers
	_ "ridge/app/controller"

	// import all module's models
	_ "ridge/app/model"

	"github.com/kataras/iris/v12"
)

// mineTypeMap file mine type
var mineTypeMap = map[string]string{
	".css": "text/css",
	".js":  "text/javascript",
}

// BuildRoute build custom controller route
func BuildRoute(app *iris.Application) {
	// index page
	app.Get("/", indexHandler)

	// favicon.ico
	app.Get("/favicon.ico", func(ctx iris.Context) {
		filePath := fmt.Sprintf("%s%s/favicon.ico",
			global.GConfig.Root, global.GConfig.App.PublicDir)
		tool.WriteFileContent(ctx, filePath)
	})

	app.Get("/logo.png", func(ctx iris.Context) {
		filePath := fmt.Sprintf("%s%s/logo.png",
			global.GConfig.Root, global.GConfig.App.PublicDir)
		tool.WriteFileContent(ctx, filePath)
	})

	// assets file
	app.Get("/assets/{anything}", staticHandler)
}

// indexHandler show index page
func indexHandler(ctx iris.Context) {
	// entry file
	// ctx.HTML("<h1>Hello World!</h1>")
	filePath := fmt.Sprintf("%s%s/index.html",
		global.GConfig.Root, global.GConfig.App.PublicDir)
	tool.WriteFileContent(ctx, filePath)
}

// staticHandler show assets file
func staticHandler(ctx iris.Context) {
	requestURL := ctx.Request().URL.Path

	// check data
	if strings.Contains(requestURL, "..") {
		indexHandler(ctx)
		return
	}

	// add head for file-mine-type
	ext := path.Ext(requestURL)
	mine, ok := mineTypeMap[ext]
	if ok {
		ctx.ContentType(mine)
	}

	filePath := fmt.Sprintf("%s%s%s", global.GConfig.Root,
		global.GConfig.App.PublicDir, requestURL)
	tool.WriteFileContent(ctx, filePath)
}
