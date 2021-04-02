/**
 * @Description: modelapp page router configuration
 */
package router

import "github.com/kataras/iris/v12"

/**
 * @Description: modelapp page router initialization
 * @param modelAppRouter: modelapp page router
 */
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
