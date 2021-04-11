/**
 * @Description: project route configuration
 */
package router

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"prometheus/api/project"
	"prometheus/model"
	"time"
)

/**
 * @Description: project api route initialization
 * @param projectAPIRouter: project api router
 */
func projectAPIInit(projectAPIRouter iris.Party) {
	// get project information
	projectAPIRouter.Get("/getProjectInfo", func(ctx iris.Context) {
		// get project list
		projectList, err := project.GetProjectList()
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
			"status":      0,
			"projectList": projectList,
		})
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}
	})

	// get project file list
	projectAPIRouter.Post("/getProjectFile", func(ctx iris.Context) {
		// get project name
		var projectJSON struct {
			ProjectName string `json:"projectName"`
		}
		// handle argument error
		if err := ctx.ReadJSON(&projectJSON); err != nil {
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
		// update current project information
		model.CurrProject = projectJSON.ProjectName

		// get file information list
		fileInfoList, err := project.GetProjectFile()
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
		ctx.StatusCode(200)
		_, err = ctx.JSON(iris.Map{
			"status":      0,
			"projectName": model.CurrProject,
			"fileList":    fileInfoList,
		})
	})

	// create project
	projectAPIRouter.Post("/createProject", func(ctx iris.Context) {
		// get project name
		var projectJSON struct {
			ProjectName string `json:"projectName"`
		}
		// handle argument error
		if err := ctx.ReadJSON(&projectJSON); err != nil {
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

		// create project dir
		_, err := project.CreateProjectDir(projectJSON.ProjectName)
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
			"message": "Create project " + projectJSON.ProjectName + " successfully",
		})
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}
	})

	// delete project
	projectAPIRouter.Post("/deleteProject", func(ctx iris.Context) {
		// get project name
		var projectJSON struct {
			ProjectName string `json:"projectName"`
		}
		// handle argument error
		if err := ctx.ReadJSON(&projectJSON); err != nil {
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

		// delete project
		_, err := project.DeleteProject(projectJSON.ProjectName)
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
			"message": "Delete project " + projectJSON.ProjectName + " successfully",
		})
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}
	})

	// create virtual env for project
	projectAPIRouter.Post("/createVirtualEnv", func(ctx iris.Context) {
		// get project name
		var projectJSON struct {
			ProjectName string `json:"projectName"`
		}
		// handle argument error
		if err := ctx.ReadJSON(&projectJSON); err != nil {
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

		// create virtual environment
		_, err := project.CreateVirtualEnv(projectJSON.ProjectName)
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
			"message": "Create virtual environment for project " + projectJSON.ProjectName + " successfully",
		})
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}
	})

	// install package requirement
	projectAPIRouter.Post("/installRequirement", func(ctx iris.Context) {
		// get project name
		var projectJSON struct {
			ProjectName string `json:"projectName"`
		}
		// handle argument error
		if err := ctx.ReadJSON(&projectJSON); err != nil {
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

		// install package requirement
		_, err := project.InstallRequirement(projectJSON.ProjectName)
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
			"message": "Install package in virtual environment for project " + projectJSON.ProjectName + " successfully",
		})
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}
	})

	// get python virtual environment package list
	projectAPIRouter.Post("/getPipList", func(ctx iris.Context) {
		// get project name
		var projectJSON struct {
			ProjectName string `json:"projectName"`
		}
		// handle argument error
		if err := ctx.ReadJSON(&projectJSON); err != nil {
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

		// get pip list
		_, err := project.GetPipList(projectJSON.ProjectName)
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
			"message": "Generate pip list for project " + projectJSON.ProjectName + " successfully",
		})
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}
	})

	// upload data file
	projectAPIRouter.Post("/uploadData", func(ctx iris.Context) {
		// get project name from cookie
		// can not get cookie when developing with Chrome locally, use other browser instead
		// reason: https://stackoverflow.com/questions/8105135/cannot-set-cookies-in-javascript
		projectName := ctx.GetCookie("projectName")
		model.CurrProject = projectName
		// save data file
		filename, err := project.UploadFile(ctx)
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
	projectAPIRouter.Post("/downloadData", func(ctx iris.Context) {
		var paramJson struct {
			Filename    string `json:"filename"`
			ProjectName string `json:"projectName"`
		}
		if err := ctx.ReadJSON(&paramJson); err != nil {
			panic(err)
		}
		golog.Info("Download data: " + paramJson.Filename + " in project: " + paramJson.ProjectName)
		err := ctx.SendFile(model.ProjectPath+"/"+paramJson.ProjectName+"/"+paramJson.Filename, paramJson.Filename)
		if err != nil {
			panic(err)
		}
	})

	// delete data file
	projectAPIRouter.Post("/deleteData", func(ctx iris.Context) {
		// get file name
		var paramJson struct {
			Filename    string `json:"filename"`
			ProjectName string `json:"projectName"`
		}
		if err := ctx.ReadJSON(&paramJson); err != nil {
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
		golog.Info("Delete data: " + paramJson.Filename + " in project: " + paramJson.ProjectName)
		// update current project information
		model.CurrProject = paramJson.ProjectName
		_, err := project.DeleteFile(paramJson.Filename)
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
