/**
 * @Description: home page route configuration
 */
package router

import "github.com/kataras/iris/v12"

/**
 * @Description: home page router initialization
 * @param homeRouter: home page router
 */
func HomeInit(homeRouter iris.Party) {
	homeRouter.Get("/", func(ctx iris.Context) {
		if err := ctx.View("index.html"); err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
		}
	})
}
