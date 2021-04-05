/**
 * @Description: datastore page route configuration
 */
package router

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	. "prometheus/api/datastore"
	"prometheus/model"
	"time"
)

/**
 * @Description: datastore page route initialization
 * @param dataStoreRouter: datastore page router
 */
func DataStoreInit(dataStoreRouter iris.Party) {
	// get datastore information
	// todo: add project func
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

	// upload data file
	dataStoreRouter.Post("/uploadData", func(ctx iris.Context) {
		// save data file
		filename, err := UploadData(ctx)
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

	// download data file
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

	// delete data file
	dataStoreRouter.Post("/deleteData", func(ctx iris.Context) {
		// get file name
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
		// delete data
		golog.Info("Delete data: " + fileJson.Filename)
		_, err := DeleteData(fileJson.Filename)
		// error in delete data
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
