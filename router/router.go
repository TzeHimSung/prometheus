package router

import (
	"fmt"
	. "prometheus/api/datastore"
	. "prometheus/api/modelstore"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func RouteInit(app *iris.Application) {
	corsConfiguration := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})

	backendRouter := app.Party("/api", corsConfiguration).AllowMethods(iris.MethodOptions)
	{
		backendRouter.Get("/getDataStoreInfo", func(ctx iris.Context) {
			_, err := ctx.JSON(iris.Map{
				"dataStoreInfo":  GetDataStoreInfo(),
				"projectList":    GetDataStoreProjectList(),
				"fileSuffixList": GetDataStoreFileSuffixList(),
			})
			if err != nil {
				panic(err)
			}
		})

		backendRouter.Post("/uploadData", func(ctx iris.Context) {
			files, n, err := ctx.UploadFormFiles("./uploads/data")
			if err != nil {
				ctx.StopWithStatus(iris.StatusInternalServerError)
				return
			}
			fmt.Printf("%d files of %d total size uploaded!\n", len(files), n)
			ctx.StatusCode(200)
			_, err = ctx.JSON(iris.Map{
				"id":     0,
				"number": len(files),
			})
			if err != nil {
				panic(err)
			}
		})

		backendRouter.Get("/getModelStoreInfo", func(ctx iris.Context) {
			_, err := ctx.JSON(iris.Map{
				"modelStoreInfo": GetModelStoreInfo(),
				"projectList":    GetModelStoreProjectList(),
				"fileSuffixList": GetModelStoreFileSuffixList(),
			})
			if err != nil {
				panic(err)
			}
		})

		backendRouter.Get("/getModelTrainingInfo", func(ctx iris.Context) {
			_, err := ctx.JSON(iris.Map{
				"info": "this is a example api",
			})
			if err != nil {
				panic(err)
			}
		})

		backendRouter.Get("/getModelAppInfo", func(ctx iris.Context) {
			_, err := ctx.JSON(iris.Map{
				"info": "this is a example api",
			})
			if err != nil {
				panic(err)
			}
		})
	}
}
