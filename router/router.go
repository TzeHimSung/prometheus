package router

import (
	"github.com/kataras/golog"
	"os"
	. "prometheus/api/datastore"
	. "prometheus/api/modelstore"
	. "prometheus/api/modeltraining"

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
			files, _, err := ctx.UploadFormFiles("./uploads/data")
			if err != nil {
				ctx.StopWithStatus(iris.StatusInternalServerError)
				return
			}
			var fileList []string
			for i := 0; i < len(files); i++ {
				fileList = append(fileList, files[i].Filename)
			}
			ctx.StatusCode(200)
			_, err = ctx.JSON(iris.Map{
				"id":       0,
				"number":   len(files),
				"filelist": fileList,
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

		dataStoreRouter.Post("/deleteData", func(ctx iris.Context) {
			var fileJson struct {
				Filename string `json:"filename"`
			}
			if err := ctx.ReadJSON(&fileJson); err != nil {
				_, err := ctx.JSON(iris.Map{
					"id":      1,
					"status":  "error",
					"message": err,
				})
				if err != nil {
					panic(err)
				}
			}
			golog.Info(fileJson.Filename)
			if err := os.Remove("./uploads/data/" + fileJson.Filename); err != nil {
				_, err := ctx.JSON(iris.Map{
					"id":      1,
					"status":  "error",
					"message": err,
				})
				if err != nil {
					panic(err)
				}
			}
			_, err := ctx.JSON(iris.Map{
				"id":      0,
				"status":  "success",
				"message": "",
			})
			if err != nil {
				panic(err)
			}
		})
	}

	modelStoreRouter := mainRouter.Party("/api")
	{
		modelStoreRouter.Get("/getModelStoreInfo", func(ctx iris.Context) {
			fileList, fileSuffixList := GetModelStoreInfo()
			projectList := GetModelStoreProjectList()
			_, err := ctx.JSON(iris.Map{
				"modelStoreInfo": fileList,
				"projectList":    projectList,
				"fileSuffixList": fileSuffixList,
			})
			if err != nil {
				panic(err)
			}
		})

		modelStoreRouter.Post("/uploadModel", func(ctx iris.Context) {
			files, _, err := ctx.UploadFormFiles("./uploads/model")
			if err != nil {
				ctx.StopWithStatus(iris.StatusInternalServerError)
				return
			}
			var fileList []string
			for i := 0; i < len(files); i++ {
				fileList = append(fileList, files[i].Filename)
			}
			ctx.StatusCode(200)
			_, err = ctx.JSON(iris.Map{
				"id":       0,
				"number":   len(files),
				"filelist": fileList,
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

		modelStoreRouter.Post("/deleteModel", func(ctx iris.Context) {
			var fileJson struct {
				Filename string `json:"filename"`
			}
			if err := ctx.ReadJSON(&fileJson); err != nil {
				_, err := ctx.JSON(iris.Map{
					"id":      1,
					"status":  "error",
					"message": err,
				})
				if err != nil {
					panic(err)
				}
			}
			golog.Info(fileJson.Filename)
			if err := os.Remove("./uploads/model/" + fileJson.Filename); err != nil {
				_, err := ctx.JSON(iris.Map{
					"id":      1,
					"status":  "error",
					"message": err,
				})
				if err != nil {
					panic(err)
				}
			}
			_, err := ctx.JSON(iris.Map{
				"id":      0,
				"status":  "success",
				"message": "",
			})
			if err != nil {
				panic(err)
			}
		})
	}

	modelTrainingRouter := mainRouter.Party("/api")
	{
		// launch model with exec
		// need to be abandoned
		modelTrainingRouter.Post("/launchtest", func(ctx iris.Context) {
			var modelJson struct {
				Modelname string `json:"modelname"`
			}
			if err := ctx.ReadJSON(&modelJson); err != nil {
				panic(err)
			}
			go LaunchModel(modelJson.Modelname)
			_, err := ctx.JSON(iris.Map{
				"status":  0,
				"message": "Model " + modelJson.Modelname + " is launched.",
			})
			if err != nil {
				panic(err)
			}
		})

		modelTrainingRouter.Get("/getModelTrainingInfo", func(ctx iris.Context) {
			_, err := ctx.JSON(iris.Map{
				"info": "this is a example api",
			})
			if err != nil {
				panic(err)
			}
		})

		modelTrainingRouter.Post("/killModel", func(ctx iris.Context) {
			var modelJson struct {
				Modelname string `json:"modelname"`
			}
			if err := ctx.ReadJSON(&modelJson); err != nil {
				panic(err)
			}
			// todo: kill model operation here
			_, err := ctx.JSON(iris.Map{
				"status":  0,
				"message": "Model " + modelJson.Modelname + " has been killed.",
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
