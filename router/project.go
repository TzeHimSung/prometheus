/**
 * @Description: project route configuration
 */
package router

import (
	"github.com/kataras/iris/v12"
	"prometheus/api/project"
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

	// todo: select project

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
			"message": "create project " + projectJSON.ProjectName + " successfully",
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
			"message": "delete project " + projectJSON.ProjectName + " successfully",
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
			"message": "create virtual environment for project " + projectJSON.ProjectName + " successfully",
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
			"message": "install package in virtual environment for project " + projectJSON.ProjectName + " successfully",
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
			"message": "generate pip list for project " + projectJSON.ProjectName + " successfully",
		})
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}
	})
}
