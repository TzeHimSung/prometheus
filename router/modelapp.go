/**
 * @Description: modelapp page router configuration
 */
package router

import (
	"github.com/kataras/iris/v12"
	. "prometheus/api/modelapp"
)

/**
 * @Description: modelapp page router initialization
 * @param modelAppRouter: modelapp page router
 */
func ModelAppInit(modelAppRouter iris.Party) {
	// test api, need to be fixed
	modelAppRouter.Get("/getModelAppInfo", func(ctx iris.Context) {
		_, err := ctx.JSON(iris.Map{
			"info": GetModelResultDir(),
		})
		if err != nil {
			panic(err)
		}
	})
}
