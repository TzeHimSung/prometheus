package router

import (
	"context"
	"github.com/kataras/iris/v12"
	"prometheus/api/database"
	. "prometheus/api/modeltraining"
	. "prometheus/model"
	"time"
)

func ModelTrainingInit(modelTrainingRouter iris.Party) {
	// get running model information
	modelTrainingRouter.Get("/getModelTrainingInfo", func(ctx iris.Context) {
		// get running model information
		runningModelList, finishedModelList := make([]FinishedModelInfo, 0), make([]FinishedModelInfo, 0)
		tmpFinishTime, _ := time.Parse(TimeFormat, "0000-00-00 00:00:00")
		for _, runningModel := range RunningModelList {
			runningModelList = append(runningModelList, FinishedModelInfo{
				Id:         runningModel.Id,
				ScriptName: runningModel.ScriptName,
				Status:     "Running",
				LaunchTime: runningModel.LaunchTime,
				FinishTime: tmpFinishTime,
			})
		}
		// get finished model information
		finishedModel, err := database.QueryFinishedModelLog()
		if err != nil {
			panic(err)
		}
		finishedModelList = append(finishedModelList, finishedModel...)
		// return response
		ctx.StatusCode(200)
		_, err = ctx.JSON(iris.Map{
			"status":            "success",
			"length":            len(runningModelList) + len(finishedModelList),
			"runningModelList":  runningModelList,
			"finishedModelList": finishedModelList,
		})
		if err != nil {
			panic(err)
		}
	})

	// launch model
	modelTrainingRouter.Post("/launchModel", func(ctx iris.Context) {
		// get model info
		modelInfo := ModelInfo{}
		if err := ctx.ReadJSON(&modelInfo); err != nil {
			panic(err)
		}

		// create model context
		modelctx, cancel := context.WithCancel(context.Background())
		// create model id
		ModelID++

		// running model record
		RunningModelList = append(RunningModelList, RunningModel{
			Id:         ModelID,
			ScriptName: modelInfo.ScriptName,
			Ctx:        modelctx,
			CancelFunc: cancel,
			LaunchTime: time.Now(),
		})

		// launch model
		go LaunchModel(modelInfo.ScriptName, ModelID, modelctx)

		// return response
		ctx.StatusCode(200)
		_, err := ctx.JSON(iris.Map{
			"status":     "success",
			"message":    "Model " + modelInfo.ScriptName + " has been launched.",
			"launchTime": time.Now().Format(TimeFormat),
		})
		if err != nil {
			panic(err)
		}
	})

	// kill running model process
	modelTrainingRouter.Post("/killModel", func(ctx iris.Context) {
		// get model info
		modelInfo := ModelInfo{}
		if err := ctx.ReadJSON(&modelInfo); err != nil {
			panic(err)
		}

		// kill model
		findModelFlag := false
		var modelIdx int = 0
		for idx, runningModel := range RunningModelList {
			if runningModel.Id == modelInfo.Id {
				modelIdx = idx
				go runningModel.CancelFunc()

				// add kill model log to database
				_, err := database.AddKilledModelLog(runningModel.Id, runningModel.ScriptName, runningModel.LaunchTime)
				if err != nil {
					panic(err)
				}

				findModelFlag = true
				break
			}
		}
		if !findModelFlag {
			ctx.StatusCode(400)
			_, err := ctx.JSON(iris.Map{
				"status":  1,
				"message": "Model " + modelInfo.ScriptName + " hasn't been launched.",
			})
			if err != nil {
				panic(err)
			}
		}

		// remove running model log
		RunningModelList = append(RunningModelList[:modelIdx], RunningModelList[modelIdx+1:]...)

		// return response
		_, err := ctx.JSON(iris.Map{
			"status":  0,
			"message": "Model " + modelInfo.ScriptName + " has been killed.",
		})
		if err != nil {
			panic(err)
		}
	})
}
