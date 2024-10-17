package autoload

import (
	"ridge/extend/autoload/custom"
	"ridge/extend/autoload/indoor"
)

// App program app
type App struct{}

// init function
type runFunc func() error

// init function list (exec order)
var funcList = []runFunc{
	indoor.AutoloadConfig,    // autoload config
	indoor.AutoloadDb,        // autoLoad db connection
	indoor.AutoMigrate,       // auto migration
	indoor.MiddleWare,        // middle ware
	indoor.AutoloadRoute,     // autoload route set
	indoor.AutoloadView,      // autoload view
	indoor.AutoloadValidator, // autoload validator
	custom.AutoloadSwagger,   // autoload swagger
	custom.AutoloadPProf,     // autoload pprof
}

// Run program app run
func (a *App) Run() error {
	// add app run func
	funcList = append(funcList, indoor.AppRun) // run iris app

	// run init func
	var err error
	for _, rFunc := range funcList {
		err = rFunc()
		if err != nil {
			return err
		}
	}

	return nil
}
