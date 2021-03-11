package main

import (
	"prometheus/router"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	router.RouteInit(app)

	if err := app.Run(iris.Addr(":8000")); err != nil {
		panic(err)
	}
}
