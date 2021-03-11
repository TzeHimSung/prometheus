package main

import (
	"prometheus/dboperation"
	"prometheus/router"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	router.RouteInit(app)

	dboperation.Dbtest()

	if err := app.Run(iris.Addr(":8000")); err != nil {
		panic(err)
	}
}
