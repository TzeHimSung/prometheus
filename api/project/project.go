// Package project
/**
 * @Description: APIs related to project
 */
package project

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"os"
	"os/exec"
	"prometheus/api/database"
	"prometheus/api/modeltraining"
	. "prometheus/model"
	"time"
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

// CreateVirtualEnv
/**
 * @Description: create virtual environment for project
 * @param projectName: project name
 * @return bool: result of creating process
 * @return error: error
 */
func CreateVirtualEnv(projectName string) (bool, error) {
	// check whether virtual environment exists or not
	_, err := os.Stat(ProjectPath + "/" + projectName + "/venv")
	// if virtual environment exists, no need to create
	if err == nil {
		return true, nil
	}
	// create project dir name
	projectDirName := ProjectPath + "/" + projectName
	// create project venv dir
	projectVEnvDirName := projectDirName + "/venv"

	// create virtual environment
	cmd := exec.Command("python", "-m", "venv", projectVEnvDirName)
	err = cmd.Run()
	if err != nil {
		return false, err
	}
	return true, nil
}

// InstallRequirement
/**
 * @Description: install virtual environment required package
 * @param projectName: project name
 * @return bool: result of installation process
 * @return error: error
 */
func InstallRequirement(projectName string) (bool, error) {
	// create project dir name
	projectDirName := ProjectPath + "/" + projectName
	// create virtual environment activate command
	venvActiveCmd := "uploads\\project\\" + projectName + "\\venv\\Scripts\\activate.bat"

	// no need to check whether requirements.txt exists or not, just return error
	// activate virtual environment
	cmd := exec.Command(venvActiveCmd, "&&",
		"pip", "install", "-r", "uploads/project/"+projectName+"/requirements.txt")
	// start installation process
	res, err := cmd.Output()
	if err != nil {
		return false, err
	}

	// create install log
	f, err := os.Create(projectDirName + "/modelLog/pip-install-" +
		time.Now().Format(modeltraining.OutputTimeFormat) + ".txt")
	defer f.Close()
	if err != nil {
		return false, err
	}
	_, err = f.Write(res)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetPipList
/**
 * @Description: get pip list info
 * @param projectName: project name
 * @return bool: result of query process
 * @return error: error
 */
func GetPipList(projectName string) (bool, error) {
	// create project dir name
	projectDirName := ProjectPath + "/" + projectName
	// create virtual environment activate command
	venvActiveCmd := "uploads\\project\\" + projectName + "\\venv\\Scripts\\activate.bat"

	// get pip list
	cmd := exec.Command(venvActiveCmd, "&&",
		"pip", "list")
	res, err := cmd.Output()
	if err != nil {
		return false, err
	}

	// create query log
	f, err := os.Create(projectDirName + "/modelLog/pip-list-" +
		time.Now().Format(modeltraining.OutputTimeFormat) + ".txt")
	defer f.Close()
	if err != nil {
		return false, err
	}
	_, err = f.Write(res)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UploadData
/**
 * @Description: save upload data file
 * @param ctx: iris context
 * @return string: file name
 * @return error: error
 */
func UploadData(ctx iris.Context) (string, error) {
	// save file to project path
	files, _, err := ctx.UploadFormFiles(ProjectPath + "/" + CurrProject)
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

// DeleteData delete data file
/**
 * @param filename: data file name
 * @return bool: result of deleting process
 * @return error: error
 */
func DeleteData(filename string) (bool, error) {
	// delete data file
	if err := os.Remove(ProjectPath + "/" + CurrProject + "/" + filename); err != nil {
		return false, err
	}
	// delete upload data log
	_, err := database.DeleteUploadDataLog(filename)
	if err != nil {
		return false, err
	}
	return true, nil
}
