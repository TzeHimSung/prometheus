package router

import (
	"context"
	"github.com/kataras/golog"
	"os"
	"prometheus/api/database"
	. "prometheus/api/datastore"
	. "prometheus/api/modelstore"
	. "prometheus/api/modeltraining"
	"prometheus/model"
	"time"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
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

			// add upload data log to database
			_, err = database.AddUploadDataLog(fileList[0])
			if err != nil {
				panic(err)
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
			golog.Info("Download data: " + fileJson.Filename)
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
			golog.Info("Delete data: " + fileJson.Filename)
			if err := os.Remove("./uploads/data/" + fileJson.Filename); err != nil {
				_, err = ctx.JSON(iris.Map{
					"id":      1,
					"status":  "error",
					"message": err,
				})
				if err != nil {
					panic(err)
				}
			}
			// delete upload data log
			_, err := database.DeleteUploadDataLog(fileJson.Filename)
			if err != nil {
				panic(err)
			}
			_, err = ctx.JSON(iris.Map{
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

			// add upload model log to database
			_, err = database.AddUploadModelLog(fileList[0])
			if err != nil {
				panic(err)
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
			// delete upload model log
			_, err := database.DeleteUploadModelLog(fileJson.Filename)
			if err != nil {
				panic(err)
			}
			_, err = ctx.JSON(iris.Map{
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

		modelTrainingRouter.Post("/launchcanceltest", func(ctx iris.Context) {
			modelctx, cancel := context.WithCancel(context.Background())

			// append running model list
			RunningModelList = append(RunningModelList, model.RunningModel{
				Id:         ModelID,
				ScriptName: "test script",
				Ctx:        &modelctx,
				CancelFunc: &cancel,
			})

			go LaunchTest(modelctx)

			time.Sleep(5 * time.Second)

			cancel()
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
