package route

import (
	"ridge/route/handler"

	"github.com/kataras/iris/v12"
)

// Build the direct route uri
func Build(app *iris.Application) {
	app.Get("/hw", func(ctx iris.Context) {
		_, _ = ctx.HTML("<h1>Hello World!</h1>")
	})

	app.Get("/hwf", handler.HelloWorld)
}
