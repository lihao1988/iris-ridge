package indoor

import (
	"fmt"

	"ridge/common/global"

	"github.com/kataras/iris/v12"
)

func AppRun() error {
	// iris app instance
	app := global.GApplication

	// set app run
	var runner iris.Runner
	gConfig := global.GConfig
	useTLS := gConfig.App.WithTLS
	addr := fmt.Sprintf("%s:%s", gConfig.App.Host, gConfig.App.Port)
	if useTLS {
		crtFile := gConfig.Root + gConfig.App.CrtFile
		keyFile := gConfig.Root + gConfig.App.KeyFile
		runner = iris.TLS(addr, crtFile, keyFile)
	} else {
		runner = iris.Addr(addr)
	}

	// app run
	return app.Run(runner)
}
