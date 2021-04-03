/**
 * @Description: APIs related to modelapp page
 */
package modelapp

import (
	"io/ioutil"
	"prometheus/model"
)

/**
 * @Description: get model result dir path
 * @return []string: model result dir list
 */
func GetModelResultDir() []string {
	filelist := make([]string, 0)
	files, _ := ioutil.ReadDir(model.ModelOutputPath)
	for _, file := range files {
		if file.IsDir() {
			filelist = append(filelist, file.Name())
		}
	}
	return filelist
}
