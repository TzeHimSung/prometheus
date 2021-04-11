// Package project
/**
 * @Description: APIs related to project
 */
package project

import (
	"github.com/kataras/golog"
	"os"
	"prometheus/api/database"
	. "prometheus/model"
)

// GetProjectList
/**
 * @Description: get project information list
 * @return []ProjectInfo: project info slice
 * @return error: error
 */
func GetProjectList() ([]Project, error) {
	// load project information from database
	projectList, err := database.QueryProjectLog()
	if err != nil {
		return nil, err
	}
	return projectList, nil
}

// GetProjectFile get project file list
/**
 * @return []FileInfo: project file list
 * @return error: error
 */
func GetProjectFile() ([]FileInfo, error) {
	fileInfoList, err := database.QueryProjectFileLog(CurrProject)
	if err != nil {
		return nil, err
	}
	return fileInfoList, nil
}

// CreateProjectDir
/**
 * @Description: create project dir
 * @param projectName: project name
 * @return bool: result of creating process
 * @return error: error
 */
func CreateProjectDir(projectName string) (bool, error) {
	// create project dir name
	projectDirName := ProjectPath + "/" + projectName
	// check whether project dir exists or not
	_, err := os.Stat(projectDirName)
	// if not exist, create dir
	if err != nil {
		// create project root dir
		err = os.Mkdir(projectDirName, 0666)
		if err != nil {
			golog.Error("Can not create project dir: " + projectName + ", please check.")
			return false, err
		}
		// create project log dir
		err = os.Mkdir(projectDirName+"/modelLog", 0666)
		if err != nil {
			golog.Error("Can not create project dir: " + projectName + ", please check.")
			return false, err
		}
	}
	// create database log
	_, err = database.AddProjectLog(projectName)
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeleteProject
/**
 * @Description: delete project dir
 * @param projectName: project name
 * @return bool: result of deleting process
 * @return error: error
 */
func DeleteProject(projectName string) (bool, error) {
	// create project dir name
	projectDirName := ProjectPath + "/" + projectName
	// delete project file
	err := os.RemoveAll(projectDirName)
	if err != nil {
		golog.Error("Can not delete project dir:" + projectName + ", please check.")
		return false, err
	}
	// delete project log from database
	_, err = database.DeleteProjectLog(projectName)
	if err != nil {
		return false, err
	}
	return true, nil
}
