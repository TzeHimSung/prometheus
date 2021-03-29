package router

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"prometheus/model"
)

var (
	ModelID          = 0
	RunningModelList = make([]model.RunningModel, 0)
)

func Hub(app *iris.Application) {
	// configure it for separation of frontend and backend development
	corsConfiguration := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:8000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	mainRouter := app.Party("/", corsConfiguration).AllowMethods(iris.MethodOptions)

	homeRouter := mainRouter.Party("/")
	HomeInit(homeRouter)

	dataStoreRouter := mainRouter.Party("/api")
	DataStoreInit(dataStoreRouter)

	modelStoreRouter := mainRouter.Party("/api")
	ModelStoreInit(modelStoreRouter)

	modelTrainingRouter := mainRouter.Party("/api")
	ModelTrainingInit(modelTrainingRouter)

	modelAppRouter := mainRouter.Party("/api")
	ModelAppInit(modelAppRouter)
}
