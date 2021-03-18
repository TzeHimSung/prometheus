package router

import (
	"fmt"
	"github.com/kataras/golog"
	. "prometheus/api/datastore"
	. "prometheus/api/modelstore"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
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
	{
		homeRouter.Get("/", func(ctx iris.Context) {
			if err := ctx.View("index.html"); err != nil {
				ctx.StopWithStatus(iris.StatusInternalServerError)
			}
		})
	}

	dataStoreRouter := mainRouter.Party("/api")
	{
		dataStoreRouter.Get("/getDataStoreInfo", func(ctx iris.Context) {
			fileList, fileSuffixList := GetDataStoreInfo()
			projectList := GetDataStoreProjectList()
			_, err := ctx.JSON(iris.Map{
				"dataStoreInfo":  fileList,
				"projectList":    projectList,
				"fileSuffixList": fileSuffixList,
			})
			if err != nil {
				panic(err)
			}
		})

		dataStoreRouter.Post("/uploadData", func(ctx iris.Context) {
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

		dataStoreRouter.Post("/downloadData", func(ctx iris.Context) {
			var fileJson struct {
				Filename string `json:"filename"`
			}
			if err := ctx.ReadJSON(&fileJson); err != nil {
				panic(err)
			}
			golog.Info(fileJson.Filename)
			if err := ctx.SendFile("./uploads/data/"+fileJson.Filename, fileJson.Filename); err != nil {
				panic(err)
			}
		})
	}

	modelStoreRouter := mainRouter.Party("/api")
	{
		modelStoreRouter.Get("/getModelStoreInfo", func(ctx iris.Context) {
			_, err := ctx.JSON(iris.Map{
				"modelStoreInfo": GetModelStoreInfo(),
				"projectList":    GetModelStoreProjectList(),
				"fileSuffixList": GetModelStoreFileSuffixList(),
			})
			if err != nil {
				panic(err)
			}
		})

		modelStoreRouter.Post("/uploadModel", func(ctx iris.Context) {
			files, n, err := ctx.UploadFormFiles("./uploads/model")
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

		modelStoreRouter.Post("/downloadModel", func(ctx iris.Context) {
			var fileJson struct {
				Filename string `json:"filename"`
			}
			if err := ctx.ReadJSON(&fileJson); err != nil {
				panic(err)
			}
			golog.Info(fileJson.Filename)
			if err := ctx.SendFile("./uploads/model/"+fileJson.Filename, fileJson.Filename); err != nil {
				panic(err)
			}
		})
	}

	modelTrainingRouter := mainRouter.Party("/api")
	{
		modelTrainingRouter.Get("/getModelTrainingInfo", func(ctx iris.Context) {
			_, err := ctx.JSON(iris.Map{
				"info": "this is a example api",
			})
			if err != nil {
				panic(err)
			}
		})
	}

	modelAppRouter := mainRouter.Party("/api")
	{
		modelAppRouter.Get("/getModelAppInfo", func(ctx iris.Context) {
			_, err := ctx.JSON(iris.Map{
				"info": "this is a example api",
			})
			if err != nil {
				panic(err)
			}
		})
	}
}
