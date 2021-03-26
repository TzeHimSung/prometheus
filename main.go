//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
package main

import (
	"flag"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	db "prometheus/api/database"
	"prometheus/router"
)

func main() {
	// parse arguments
	flag.Parse()

	app := iris.New()

	// database init
	_, err := db.InitDatabase()
	if err != nil {
		panic(err)
	}

	// router init
	router.Hub(app)

	// launch server
	golog.SetLevel("debug")
	golog.Info("prometheus is launching...")

	app.RegisterView(iris.HTML("dist", ".html"))
	app.HandleDir("/static", "dist/static")

	if err := app.Run(iris.Addr(":8000")); err != nil {
		panic(err)
	}
}
