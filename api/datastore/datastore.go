package datastore

import (
	. "prometheus/model"
)

func GetDataStoreProjectList() []ProjectInfo {
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

func GetDataStoreFileSuffixList() []FileSuffixInfo {
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

func GetDataStoreInfo() []DataStoreInfo {
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
