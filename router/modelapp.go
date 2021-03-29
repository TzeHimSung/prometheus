package router

import "github.com/kataras/iris/v12"

func ModelAppInit(modelAppRouter iris.Party) {
	modelAppRouter.Get("/getModelAppInfo", func(ctx iris.Context) {
		_, err := ctx.JSON(iris.Map{
			"info": "this is a example api",
		})
		if err != nil {
			panic(err)
		}
	})
}
