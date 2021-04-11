/**
 * @Description: route configuration
 */
package router

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

/**
 * @Description: route initialization
 * @param app: iris application
 */
func Hub(app *iris.Application) {
	// configure it for separation of frontend and backend development
	corsConfiguration := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:8000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// config main router
	mainRouter := app.Party("/", corsConfiguration).AllowMethods(iris.MethodOptions)

	// project API router
	projectAPIRouter := mainRouter.Party("/api")
	projectAPIInit(projectAPIRouter)

	// home page router init
	homeRouter := mainRouter.Party("/")
	HomeInit(homeRouter)

	// modeltraining page router init
	modelTrainingRouter := mainRouter.Party("/api")
	ModelTrainingInit(modelTrainingRouter)

	// modelapp page router init
	modelAppRouter := mainRouter.Party("/api")
	ModelAppInit(modelAppRouter)
}
