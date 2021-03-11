package modelstore

import (
	. "prometheus/datastructure"
)

func GetModelStoreProjectList() []ProjectInfo {
	return []ProjectInfo{
		{
			Id:   0,
			Name: "Project AA",
		},
		{
			Id:   1,
			Name: "Project BB",
		},
		{
			Id:   2,
			Name: "Project CC",
		},
	}
}

func GetModelStoreFileSuffixList() []FileSuffixInfo {
	return []FileSuffixInfo{
		{
			Id:   0,
			Name: "py",
		},
		{
			Id:   1,
			Name: "other",
		},
	}
}

func GetModelStoreInfo() []ModelStoreInfo {
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
