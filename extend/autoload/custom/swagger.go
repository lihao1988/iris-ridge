package custom

import (
	"fmt"

	"ridge/common/global"

	// register swag
	_ "ridge/docs"

	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
)

// AutoloadSwagger autoload swagger
func AutoloadSwagger() error {
	// open swagger
	if !global.GConfig.App.WithSwagger {
		return nil
	}

	// app host
	var swaggerUrl string
	addr := fmt.Sprintf("%s:%s", global.GConfig.App.Host, global.GConfig.App.Port)
	useTLS := global.GConfig.App.WithTLS
	if useTLS {
		swaggerUrl = fmt.Sprintf("https://%s/swagger/swagger.json", addr)
	} else {
		swaggerUrl = fmt.Sprintf("http://%s/swagger/swagger.json", addr)
	}

	// Configure the swagger UI page.
	swaggerUI := swagger.Handler(swaggerFiles.Handler,
		swagger.URL(swaggerUrl),
		swagger.DeepLinking(true),
		swagger.Prefix("/swagger"),
	)

	// Register on http://{addr}/swagger
	app := global.GApplication // iris app
	app.Get("/swagger", swaggerUI)
	// And http://{addr}/swagger/index.html, *.js, *.css and e.t.c.
	app.Get("/swagger/{any:path}", swaggerUI)

	return nil
}
