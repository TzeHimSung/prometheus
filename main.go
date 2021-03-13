//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
package main

import (
	"prometheus/dboperation"
	"prometheus/router"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	router.Hub(app)

	dboperation.DbTest()

	app.RegisterView(iris.HTML("dist", ".html"))
	app.HandleDir("/static", "dist/static")

	if err := app.Run(iris.Addr(":8000")); err != nil {
		panic(err)
	}
}
