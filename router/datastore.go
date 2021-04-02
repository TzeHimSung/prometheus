/**
 * @Description: datastore page route configuration
 */
package router

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"os"
	"prometheus/api/database"
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
	// attention: project list may be removed
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
	// todo: rewrite this api, move part function to API package
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
			"id":         0,
			"number":     len(files),
			"filelist":   fileList,
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
	// todo: rewrite this api, move part function to API package
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
