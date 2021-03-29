package router

import (
	"context"
	"github.com/kataras/iris/v12"
	. "prometheus/api/modeltraining"
	"prometheus/model"
	"time"
)

func ModelTrainingInit(modelTrainingRouter iris.Party) {
	// launch model with exec
	// need to be abandoned
	modelTrainingRouter.Post("/launchtest", func(ctx iris.Context) {
		var modelJson struct {
			Modelname string `json:"modelname"`
		}
		if err := ctx.ReadJSON(&modelJson); err != nil {
			panic(err)
		}

		go LaunchModel(modelJson.Modelname)
		_, err := ctx.JSON(iris.Map{
			"status":  0,
			"message": "Model " + modelJson.Modelname + " is launched.",
		})
		if err != nil {
			panic(err)
		}
	})

	modelTrainingRouter.Post("/launchcanceltest", func(ctx iris.Context) {
		modelctx, cancel := context.WithCancel(context.Background())

		// append running model list
		RunningModelList = append(RunningModelList, model.RunningModel{
			Id:         ModelID,
			ScriptName: "test script",
			Ctx:        &modelctx,
			CancelFunc: &cancel,
		})

		go LaunchTest(modelctx)

		time.Sleep(5 * time.Second)

		cancel()
	})

	modelTrainingRouter.Get("/getModelTrainingInfo", func(ctx iris.Context) {
		_, err := ctx.JSON(iris.Map{
			"info": "this is a example api",
		})
		if err != nil {
			panic(err)
		}
	})

	modelTrainingRouter.Post("/killModel", func(ctx iris.Context) {
		var modelJson struct {
			Modelname string `json:"modelname"`
		}
		if err := ctx.ReadJSON(&modelJson); err != nil {
			panic(err)
		}
		// todo: kill model operation here
		_, err := ctx.JSON(iris.Map{
			"status":  0,
			"message": "Model " + modelJson.Modelname + " has been killed.",
		})
		if err != nil {
			panic(err)
		}
	})
}
