/**
 * @Description: APIs related to modelstore page
 */
package modelstore

import (
	"github.com/kataras/iris/v12"
	"os"
	"prometheus/api/database"
	. "prometheus/model"
	"strings"
)

/**
 * @Description: get project information, this API may be need to be removed
 * @return []ProjectInfo: project information slice
 */
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

/**
 * @Description: get model store file information
 * @return []ModelStoreInfo: model store file slice
 * @return []FileSuffixInfo: model store file suffix slice
 */
func GetModelStoreInfo() ([]ModelStoreInfo, []FileSuffixInfo) {
	// load file list data from database
	fileList, err := database.QueryUploadModelLog()
	if err != nil {
		panic(err)
	}

	// get file suffix
	fileSuffixList := make([]string, 0, 10)
	for _, file := range fileList {
		fileSuffixList = append(fileSuffixList, strings.Split(file.FileName, ".")[len(strings.Split(file.FileName, "."))-1])
	}

	// remove duplicated file suffix
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

	// return file list and file suffix list
	return fileList, resultFileSuffixList
}

/**
 * @Description: save upload model file
 * @param ctx: iris context
 * @return string: file name
 * @return error: error
 */
func UploadModel(ctx iris.Context) (string, error) {
	// save file
	files, _, err := ctx.UploadFormFiles(ModelPath)
	if err != nil {
		return "", err
	}

	// add upload model log to database
	// here must be files[0] because its multipart file upload
	_, err = database.AddUploadModelLog(files[0].Filename)
	if err != nil {
		panic(err)
	}
	return files[0].Filename, nil
}

/**
 * @Description: delete model file
 * @param filename: model file name
 * @return bool: result of deleting process
 * @return error: error
 */
func DeteleModel(filename string) (bool, error) {
	// delete model file
	if err := os.Remove(ModelPath + "/" + filename); err != nil {
		return false, err
	}
	// delete upload model log
	_, err := database.DeleteUploadModelLog(filename)
	if err != nil {
		return false, err
	}
	return true, nil
}
