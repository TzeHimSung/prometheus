//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
package main

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	db "prometheus/api/database"
	"prometheus/router"
)

const (
	ListenPort = 8000
)

func main() {
	// get iris app object
	app := iris.New()

	// database init
	_, err := db.InitDatabase()
	if err != nil {
		panic(err)
	}

	// router init
	router.Hub(app)

	// golog configuration
	golog.SetLevel("debug")
	golog.Info("prometheus is launching...")

	// register Vue dist and static resource path
	app.RegisterView(iris.HTML("dist", ".html"))
	app.HandleDir("/static", "dist/static")

	// launch server
	if err := app.Run(iris.Addr(fmt.Sprintf(":%d", ListenPort))); err != nil {
		panic(err)
	}
}
