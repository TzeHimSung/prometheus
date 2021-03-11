/*
 * @Author: JHSeng
 * @Date: 2021-03-11 16:17:12
 * @LastEditTime: 2021-03-11 16:19:43
 * @LastEditors: Please set LastEditors
 * @Description: main logic for prometheus
 * @FilePath: \prometheus\main.go
 */

/*
 * @Author: JHSeng
 * @Date: 2021-03-02 10:55:59
 * @LastEditTime: 2021-03-11 16:11:12
 * @LastEditors: Please set LastEditors
 * @Description: main logic of project prometheus
 * @FilePath: \myapp\main.go
 */

package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

type ProjectInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type FileSuffixInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DataStoreInfo struct {
	FileName   string `json:"fileName"`
	Source     string `json:"source"`
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
}

type ModelStoreInfo struct {
	FileName   string `json:"fileName"`
	Source     string `json:"source"`
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
}

func getDataStoreProjectList() []ProjectInfo {
	return []ProjectInfo{
		{
			Id:   0,
			Name: "Project A",
		},
		{
			Id:   1,
			Name: "Project B",
		},
		{
			Id:   2,
			Name: "Project C",
		},
	}
}

func getDataStoreFileSuffixList() []FileSuffixInfo {
	return []FileSuffixInfo{
		{
			Id:   0,
			Name: "csv",
		},
		{
			Id:   1,
			Name: "txt",
		},
		{
			Id:   2,
			Name: "json",
		},
		{
			Id:   3,
			Name: "other",
		},
	}
}

/**
 * @description: get stored data information
 * @param none
 * @return []DataStoreInfo
 */
func getDataStoreInfo() []DataStoreInfo {
	return []DataStoreInfo{
		{
			FileName:   "fileName1",
			Source:     "user upload",
			Status:     "上传中",
			CreateTime: "2021-03-08 00:00:00",
		},
		{
			FileName:   "fileName2",
			Source:     "user upload",
			Status:     "已上传",
			CreateTime: "2021-03-08 00:00:00",
		},
		{
			FileName:   "fileName3",
			Source:     "user upload",
			Status:     "上传中",
			CreateTime: "2021-03-08 00:00:00",
		},
		{
			FileName:   "fileName4",
			Source:     "user upload",
			Status:     "已上传",
			CreateTime: "2021-03-08 00:00:00",
		},
		{
			FileName:   "fileName5",
			Source:     "user upload",
			Status:     "上传中",
			CreateTime: "2021-03-08 00:00:00",
		},
	}
}

/**
 * @description: get stored model information
 * @param none
 * @return []ModelStoreInfo
 */
func getModelStoreInfo() []ModelStoreInfo {
	return []ModelStoreInfo{
		{
			FileName:   "modelName1",
			Source:     "user upload",
			Status:     "上传中",
			CreateTime: "2021-03-08 00:00:00",
		},
		{
			FileName:   "modelName2",
			Source:     "user upload",
			Status:     "已上传",
			CreateTime: "2021-03-08 00:00:00",
		},
		{
			FileName:   "modelName3",
			Source:     "user upload",
			Status:     "上传中",
			CreateTime: "2021-03-08 00:00:00",
		},
		{
			FileName:   "modelName4",
			Source:     "user upload",
			Status:     "已上传",
			CreateTime: "2021-03-08 00:00:00",
		},
		{
			FileName:   "modelName5",
			Source:     "user upload",
			Status:     "上传中",
			CreateTime: "2021-03-08 00:00:00",
		},
	}
}

func main() {
	app := iris.New()

	corsConfiguration := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})

	backendRouter := app.Party("/api", corsConfiguration).AllowMethods(iris.MethodOptions)
	{
		backendRouter.Get("/getDataStoreInfo", func(ctx iris.Context) {
			ctx.JSON(iris.Map{
				"dataStoreInfo":  getDataStoreInfo(),
				"projectList":    getDataStoreProjectList(),
				"fileSuffixList": getDataStoreFileSuffixList(),
			})
		})

		backendRouter.Get("/getModelStoreInfo", func(ctx iris.Context) {
			ctx.JSON(iris.Map{
				"modelStoreInfo": getModelStoreInfo(),
			})
		})

		backendRouter.Get("/getModelTrainingInfo", func(ctx iris.Context) {
			ctx.JSON(iris.Map{
				"info": "this is a example api",
			})
		})

		backendRouter.Get("/getModelAppInfo", func(ctx iris.Context) {
			ctx.JSON(iris.Map{
				"info": "this is a example api",
			})
		})

		backendRouter.Post("/testpost", func(ctx iris.Context) {
			ctx.WriteString("Post method test is success!")
		})
	}

	app.Run(iris.Addr(":8000"))
}
