/**
 * @Description: APIs related to datastore page
 */
package datastore

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

/**
 * @Description: get data store file information
 * @return []DataStoreInfo: data store file slice
 * @return []FileSuffixInfo: data store file suffix slice
 */
func GetDataStoreInfo() ([]DataStoreInfo, []FileSuffixInfo) {
	// load file list data from database
	fileList, err := database.QueryUploadDataLog()
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
 * @Description: save upload data
 * @param ctx: iris context
 * @return string: filename
 * @return error: error
 */
func UploadData(ctx iris.Context) (string, error) {
	// save file
	files, _, err := ctx.UploadFormFiles(DataPath)
	if err != nil {
		return "", err
	}

	// add upload data log to database
	// here must be files[0] because its multipart file upload
	_, err = database.AddUploadDataLog(files[0].Filename)
	if err != nil {
		panic(err)
	}
	return files[0].Filename, nil
}

/**
 * @Description: delete data file
 * @param filename: data file name
 * @return bool: result of deleting process
 * @return error: error
 */
func DeleteData(filename string) (bool, error) {
	// delete data file
	if err := os.Remove(DataPath + "/" + filename); err != nil {
		return false, err
	}
	// delete upload data log
	_, err := database.DeleteUploadDataLog(filename)
	if err != nil {
		return false, err
	}
	return true, nil
}
