package handler

import (
	"github.com/kataras/iris/v12"
)

// HelloWorld the first step of all
func HelloWorld(ctx iris.Context) {
	ctx.HTML("<h1>Hello World function!</h1>")
}
