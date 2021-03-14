//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
package main

import (
	"flag"
	"prometheus/dboperation"
	"prometheus/router"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
)

var initDB = flag.Bool("initdb", false, "初始化数据库")
var runServer = flag.Bool("runserver", false, "启动服务")

func main() {
	// parse arguments
	flag.Parse()

	app := iris.New()

	router.Hub(app)

	if *initDB {
		dboperation.DbTest()
	}

	if *runServer {
		golog.SetLevel("debug")
		golog.Info("prometheus is launching...")

		app.RegisterView(iris.HTML("dist", ".html"))
		app.HandleDir("/static", "dist/static")

		if err := app.Run(iris.Addr(":8000")); err != nil {
			panic(err)
		}
	}
}
