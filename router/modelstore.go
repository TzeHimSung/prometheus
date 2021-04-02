/**
 * @Description: modelstore page router configuration
 */
package router

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"os"
	"prometheus/api/database"
	. "prometheus/api/modelstore"
	"prometheus/model"
	"time"
)

/**
 * @Description: modelstore page router initialization
 * @param modelStoreRouter: modelstore page router
 */
func ModelStoreInit(modelStoreRouter iris.Party) {
	// get modelstore information
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

	// upload model file
	// todo: rewrite this api, move part function to API package
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
			"id":         0,
			"number":     len(files),
			"filelist":   fileList,
			"createTime": time.Now().Format(model.TimeFormat),
		})
		if err != nil {
			panic(err)
		}
	})

	// download model file
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

	// delete model file
	// todo: rewrite this api, move part function to API package
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
