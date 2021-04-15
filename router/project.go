// Package router
/**
 * @Description: project route configuration
 */
package router

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"log"
	"prometheus/api/database"
	. "prometheus/api/project"
	. "prometheus/model"
	"time"
)

// projectAPIInit project api route initialization
/**
 * @param projectAPIRouter: project api router
 */
func projectAPIInit(projectAPIRouter iris.Party) {
	// get project information
	projectAPIRouter.Get("/getProjectInfo", func(ctx iris.Context) {
		// get project list
		projectList, err := GetProjectList()
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
		CurrProject = projectJSON.ProjectName

		// get file information list
		fileInfoList, err := GetProjectFile()
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
			"projectName": CurrProject,
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
		_, err := CreateProjectDir(projectJSON.ProjectName)
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
		_, err := DeleteProject(projectJSON.ProjectName)
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
		_, err := CreateVirtualEnv(projectJSON.ProjectName)
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
		_, err := InstallRequirement(projectJSON.ProjectName)
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
		_, err := GetPipList(projectJSON.ProjectName)
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

	// upload project file
	projectAPIRouter.Post("/uploadFile", func(ctx iris.Context) {
		// get project name from cookie
		// can not get cookie when developing with Chrome locally, use other browser instead
		// reason: https://stackoverflow.com/questions/8105135/cannot-set-cookies-in-javascript
		projectName := ctx.GetCookie("projectName")
		CurrProject = projectName
		// save data file
		filename, err := UploadFile(ctx)
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
			"createTime": time.Now().Format(TimeFormat),
		})
		if err != nil {
			panic(err)
		}
	})

	// download project file
	projectAPIRouter.Post("/downloadFile", func(ctx iris.Context) {
		var paramJSON struct {
			Filename    string `json:"filename"`
			ProjectName string `json:"projectName"`
		}
		if err := ctx.ReadJSON(&paramJSON); err != nil {
			panic(err)
		}
		golog.Info("Download data: " + paramJSON.Filename + " in project: " + paramJSON.ProjectName)
		err := ctx.SendFile(ProjectPath+"/"+paramJSON.ProjectName+"/"+paramJSON.Filename, paramJSON.Filename)
		if err != nil {
			panic(err)
		}
	})

	// delete project file
	projectAPIRouter.Post("/deleteFile", func(ctx iris.Context) {
		// get file and project name
		var paramJSON struct {
			Filename    string `json:"filename"`
			ProjectName string `json:"projectName"`
		}
		if err := ctx.ReadJSON(&paramJSON); err != nil {
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
		golog.Info("Delete data: " + paramJSON.Filename + " in project: " + paramJSON.ProjectName)
		// update current project information
		CurrProject = paramJSON.ProjectName
		_, err := DeleteFile(paramJSON.Filename)
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

	// launch project
	projectAPIRouter.Post("/launchProject", func(ctx iris.Context) {
		// get project name
		var paramJSON struct {
			ProjectName string `json:"projectName"`
		}
		if err := ctx.ReadJSON(&paramJSON); err != nil {
			_, err := ctx.JSON(iris.Map{
				"status":  "error",
				"message": err,
			})
			if err != nil {
				panic(err)
			}
		}

		// update current project
		CurrProject = paramJSON.ProjectName
		// add project id
		ProjectID++
		// make quit channel
		quitChan := make(chan int)

		// launch model
		go LaunchProject(paramJSON.ProjectName, ProjectID, quitChan)

		// return response
		ctx.StatusCode(200)
		_, err := ctx.JSON(iris.Map{
			"status":     "success",
			"message":    "Project " + paramJSON.ProjectName + " has been launched.",
			"launchTime": time.Now().Format(TimeFormat),
		})
		if err != nil {
			panic(err)
		}
	})

	// get running project information
	projectAPIRouter.Get("/getRunningProjectInfo", func(ctx iris.Context) {
		runningProjectList := make([]FinishedProjectInfo, 0)
		// get running project information
		tmpFinishTime, _ := time.Parse(TimeFormat, "0000-00-00 00:00:00")
		for _, runningProject := range RunningProjectList {
			runningProjectList = append(runningProjectList, FinishedProjectInfo{
				Id:          runningProject.Id,
				Pid:         runningProject.Pid,
				ProjectName: runningProject.ProjectName,
				Status:      "Running",
				LaunchTime:  runningProject.LaunchTime,
				FinishTime:  tmpFinishTime,
			})
		}
		// get finished project information
		finishedProjectList, err := database.QueryFinishedProjectLog()
		if err != nil {
			panic(err)
		}
		// return response
		ctx.StatusCode(200)
		_, err = ctx.JSON(iris.Map{
			"status":              "success",
			"length":              len(runningProjectList) + len(finishedProjectList),
			"runningProjectList":  runningProjectList,
			"finishedProjectList": finishedProjectList,
		})
		if err != nil {
			panic(err)
		}
	})

	// kill running project
	projectAPIRouter.Post("/killProject", func(ctx iris.Context) {
		// get project id and name
		var paramJSON struct {
			Id          int    `json:"id"`
			ProjectName string `json:"projectName"`
		}
		if err := ctx.ReadJSON(&paramJSON); err != nil {
			_, err := ctx.JSON(iris.Map{
				"id":      1,
				"status":  "error",
				"message": err,
			})
			if err != nil {
				panic(err)
			}
		}

		// kill model
		var findProjectFlag, projectIdx = false, 0
		for idx, runningProject := range RunningProjectList {
			if runningProject.Id == paramJSON.Id {
				// mark project id
				projectIdx = idx

				// cancel running project goroutine
				err := KillProcessWithPid(runningProject.Pid)
				if err != nil {
					log.Fatal(err)
				}

				// close quit channel
				close(runningProject.QuitChan)

				// add kill project log to database
				_, err = database.AddKilledProjectLog(
					runningProject.Id,
					runningProject.Pid,
					runningProject.ProjectName,
					runningProject.LaunchTime)

				if err != nil {
					panic(err)
				}
				findProjectFlag = true
				break
			}
		}

		// if not find project
		if !findProjectFlag {
			ctx.StatusCode(400)
			_, err := ctx.JSON(iris.Map{
				"status":  "error",
				"message": "Project " + paramJSON.ProjectName + " hasn't been launched.",
			})
			if err != nil {
				panic(err)
			}
		}

		// mark project launch time
		var projectLaunchTime = RunningProjectList[projectIdx].LaunchTime

		// remove running project
		RunningProjectList = append(RunningProjectList[:projectIdx], RunningProjectList[projectIdx+1:]...)

		// return response
		_, err := ctx.JSON(iris.Map{
			"status":            "success",
			"message":           "Project " + paramJSON.ProjectName + " has been killed.",
			"projectName":       paramJSON.ProjectName,
			"projectStatus":     "Killed",
			"projectLaunchTime": projectLaunchTime.Format(TimeFormat),
		})
		if err != nil {
			panic(err)
		}
	})
}
