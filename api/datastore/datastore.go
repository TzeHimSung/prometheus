package datastore

import (
	"io/ioutil"
	. "prometheus/model"
	"strings"
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

func GetDataStoreInfo() ([]DataStoreInfo, []FileSuffixInfo) {
	fileList, fileSuffixList := make([]DataStoreInfo, 0, 10), make([]FileSuffixInfo, 0, 10)

	dir, err := ioutil.ReadDir("./uploads/data")
	if err != nil {
		panic(err)
	}

	suffixCount := 0
	for _, file := range dir {
		if file.IsDir() {
			continue
		}
		fileList = append(fileList, DataStoreInfo{
			FileName:   file.Name(),
			Source:     "user upload",
			Status:     "已上传",
			CreateTime: "2021-03-08 00:00:00",
		})
		fileSuffixList = append(fileSuffixList, FileSuffixInfo{
			Id:   suffixCount,
			Name: strings.Split(file.Name(), ".")[len(strings.Split(file.Name(), "."))-1],
		})
		suffixCount++
	}

	resultFileSuffixList := make([]FileSuffixInfo, 0, len(fileSuffixList))
	tempMap := map[FileSuffixInfo]struct{}{}

	for _, item := range fileSuffixList {
		if _, ok := tempMap[item]; !ok {
			tempMap[item] = struct{}{}
			resultFileSuffixList = append(resultFileSuffixList, item)
		}
	}

	return fileList, resultFileSuffixList
}
