package custom

import (
	"ridge/common/global"

	"github.com/kataras/iris/v12/middleware/pprof"
)

// AutoloadPProf autoload pprof
func AutoloadPProf() error {
	// open pprof
	if !global.GConfig.App.WithPProf {
		return nil
	}

	// iris app
	app := global.GApplication
	app.HandleMany("GET", "/debug/pprof /debug/pprof/{action:path}", pprof.New())

	return nil
}
