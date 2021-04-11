// Package project
/**
 * @Description: APIs related to project
 */
package project

import (
	"github.com/kataras/iris/v12"
	"os"
	"prometheus/api/database"
	. "prometheus/model"
)

// UploadFile
/**
 * @Description: save upload project file
 * @param ctx: iris context
 * @return string: file name
 * @return error: error
 */
func UploadFile(ctx iris.Context) (string, error) {
	// save file to project path
	files, _, err := ctx.UploadFormFiles(ProjectPath + "/" + CurrProject)
	if err != nil {
		return "", err
	}

	// add upload data log to database
	// here must be files[0] because its multipart file upload
	_, err = database.AddUploadFileLog(files[0].Filename)
	if err != nil {
		panic(err)
	}
	return files[0].Filename, nil
}

// DeleteFile delete project file
/**
 * @param filename: project file name
 * @return bool: result of deleting process
 * @return error: error
 */
func DeleteFile(filename string) (bool, error) {
	// delete data file
	if err := os.Remove(ProjectPath + "/" + CurrProject + "/" + filename); err != nil {
		return false, err
	}
	// delete upload data log
	_, err := database.DeleteUploadFileLog(filename)
	if err != nil {
		return false, err
	}
	return true, nil
}
