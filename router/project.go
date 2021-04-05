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
		ctx.StatusCode(200)
		_, err := ctx.JSON(iris.Map{
			"status": 0,
		})
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}
	})

	// create project
	projectAPIRouter.Post("/createProject", func(ctx iris.Context) {
		// get project name
		var projectJSON struct {
			ProjectName string `json:"projectName"`
		}
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
			"status": 0,
		})
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}
	})

	// delete project
	projectAPIRouter.Post("/deleteProject", func(ctx iris.Context) {
		ctx.StatusCode(200)
		_, err := ctx.JSON(iris.Map{
			"status": 0,
		})
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}
	})

	// create virtual env for project
	projectAPIRouter.Post("/createVirtualEnv", func(ctx iris.Context) {
		ctx.StatusCode(200)
		_, err := ctx.JSON(iris.Map{
			"status": 0,
		})
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}
	})
}
