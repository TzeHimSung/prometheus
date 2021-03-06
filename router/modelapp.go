/**
 * @Description: modelapp page router configuration
 */
package router

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"path/filepath"
	. "prometheus/api/modelapp"
	"prometheus/api/project"
)

// ModelAppInit modelapp page router initialization
/**
 * @param modelAppRouter: modelapp page router
 */
func ModelAppInit(modelAppRouter iris.Party) {
	// get project result
	modelAppRouter.Get("/getProjectResult", func(ctx iris.Context) {
		// get project output dir
		projectResultList, err := GetProjectResultDir()
		if err != nil {
			ctx.StatusCode(500)
			_, err := ctx.JSON(iris.Map{
				"status":  "error",
				"message": err,
			})
			if err != nil {
				ctx.StopWithStatus(iris.StatusInternalServerError)
				return
			}
		}

		// return response
		ctx.StatusCode(200)
		_, err = ctx.JSON(iris.Map{
			"status":            0,
			"projectResultList": projectResultList,
		})
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}
	})

	// read specific project result
	modelAppRouter.Post("/loadProjectResult", func(ctx iris.Context) {
		// get project output dir name
		var paramJSON struct {
			OutputDirName string `json:"outputDirName"`
		}
		if err := ctx.ReadJSON(&paramJSON); err != nil {
			ctx.StatusCode(400)
			_, err := ctx.JSON(iris.Map{
				"status":  "error",
				"message": err,
			})
			if err != nil {
				ctx.StopWithStatus(iris.StatusInternalServerError)
				return
			}
		}

		// read project result
		res, err := LoadProjectResult(paramJSON.OutputDirName)
		if err != nil {
			ctx.StatusCode(500)
			_, err := ctx.JSON(iris.Map{
				"status":  "error",
				"message": err,
			})
			if err != nil {
				ctx.StopWithStatus(iris.StatusInternalServerError)
				return
			}
		}

		// return response
		ctx.StatusCode(200)
		_, err = ctx.JSON(iris.Map{
			"status":  0,
			"content": res,
		})
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}
	})

	modelAppRouter.Post("/downloadResult", func(ctx iris.Context) {
		var paramJSON struct {
			ResultDirName string `json:"resultDirName"`
		}
		if err := ctx.ReadJSON(&paramJSON); err != nil {
			panic(err)
		}
		golog.Info("Download project result: " + paramJSON.ResultDirName)
		err := ctx.SendFile(filepath.Join(project.OutputRootPath, paramJSON.ResultDirName, "output.json"),
			"output.json")
		if err != nil {
			panic(err)
		}
	})
}
