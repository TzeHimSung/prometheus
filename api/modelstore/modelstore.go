package modelstore

import (
	"io/ioutil"
	. "prometheus/model"
	"strings"
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

func GetModelStoreInfo() ([]ModelStoreInfo, []FileSuffixInfo) {
	fileList, fileSuffixList := make([]ModelStoreInfo, 0, 10), make([]string, 0, 10)

	dir, err := ioutil.ReadDir("./uploads/model")
	if err != nil {
		panic(err)
	}

	for _, file := range dir {
		if file.IsDir() {
			continue
		}
		fileList = append(fileList, ModelStoreInfo{
			FileName:   file.Name(),
			Source:     "user upload",
			Status:     "已上传",
			CreateTime: "2021-03-08 00:00:00",
		})
		fileSuffixList = append(fileSuffixList, strings.Split(file.Name(), ".")[len(strings.Split(file.Name(), "."))-1])
	}

	resultFileSuffixList := make([]FileSuffixInfo, 0, len(fileSuffixList))
	tempMap := map[string]struct{}{}

	suffixCount := 0
	for _, item := range fileSuffixList {
		if _, ok := tempMap[item]; !ok {
			tempMap[item] = struct{}{}
			resultFileSuffixList = append(resultFileSuffixList, FileSuffixInfo{
				Id:   suffixCount,
				Name: item,
			})
			suffixCount++
		}
	}

	return fileList, resultFileSuffixList
}
