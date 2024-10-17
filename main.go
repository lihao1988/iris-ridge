package main

import (
	"ridge/common/global"
	"ridge/extend/autoload"
	"ridge/tool"

	"github.com/kataras/iris/v12"
)

// main the entry program function
func main() {
	// program root path
	global.GConfig.Root = tool.GetExecFilePath()

	// iris app (global)
	global.GApplication = iris.New()

	// app run
	err := new(autoload.App).Run()
	if err != nil {
		panic(err)
	}
}
