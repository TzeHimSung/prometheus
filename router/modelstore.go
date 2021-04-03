/**
 * @Description: modelstore page router configuration
 */
package router

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
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
	modelStoreRouter.Post("/uploadModel", func(ctx iris.Context) {
		// save model file
		filename, err := UploadModel(ctx)
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}

		// return response
		ctx.StatusCode(200)
		_, err = ctx.JSON(iris.Map{
			"id":         0,
			"status":     "success",
			"filename":   filename,
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
		// delete model
		golog.Info("Delete model: " + fileJson.Filename)
		_, err := DeteleModel(fileJson.Filename)
		if err != nil {
			_, err = ctx.JSON(iris.Map{
				"id":      1,
				"status":  "error",
				"message": err,
			})
			if err != nil {
				panic(err)
			}
		}
		// handle success
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
