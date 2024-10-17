package indoor

import (
	"fmt"
	"ridge/common/global"

	"github.com/kataras/iris/v12"
)

// AutoloadView autoload view set
func AutoloadView() error {
	// iris app
	app := global.GApplication

	// view template
	viewDir := fmt.Sprintf("%s%s", global.GConfig.Root, global.GConfig.App.ViewDir)
	tmpl := iris.HTML(viewDir, ".html")
	tmpl.Reload(true) // reload templates on each request (development mode)

	app.RegisterView(tmpl)

	return nil
}
