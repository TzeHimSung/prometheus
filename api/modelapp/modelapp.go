/**
 * @Description: APIs related to modelapp page
 */
package modelapp

import (
	"io/ioutil"
	"prometheus/model"
)

/**
 * @Description: get model result dir list
 * @return []string: model result dir list
 */
func GetModelResultDir() ([]string, error) {
	dirList := make([]string, 0)
	files, err := ioutil.ReadDir(model.ModelOutputPath)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			dirList = append(dirList, file.Name())
		}
	}
	return dirList, nil
}

func LoadProjectResult(dirName string) string {
	return ""
}
