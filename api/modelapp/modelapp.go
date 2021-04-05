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
func GetModelResultDir() []string {
	dirList := make([]string, 0)
	files, _ := ioutil.ReadDir(model.ModelOutputPath)
	for _, file := range files {
		if file.IsDir() {
			dirList = append(dirList, file.Name())
		}
	}
	return dirList
}
