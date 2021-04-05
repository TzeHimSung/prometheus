/**
 * @Description: APIs related to project
 */
package project

import (
	"github.com/kataras/golog"
	"os"
	"os/exec"
	. "prometheus/model"
)

/**
 * @Description: create project dir
 * @param projectName: project name
 * @return bool: sign of creating process
 * @return error: error
 */
func CreateProjectDir(projectName string) (bool, error) {
	// create project dir name
	projectDirName := ProjectPath + "/" + projectName
	// check whether project dir exists or not
	_, err := os.Stat(projectDirName)
	// if not exist, create dir
	if err != nil {
		err = os.Mkdir(projectDirName, 0666)
		if err != nil {
			golog.Error("Can not create project dir: " + projectName + ", please check.")
			return false, err
		}
	}
	return true, nil
}

/**
 * @Description: delete project dir
 * @param projectName: project name
 * @return bool: sign of deleting process
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
	return true, nil
}

/**
 * @Description: create virtual environment for project
 * @param projectName: project name
 * @return bool: sign of creating process
 * @return error: error
 */
func CreateVirtualEnv(projectName string) (bool, error) {
	// create project dir name
	projectDirName := ProjectPath + "/" + projectName
	// create project venv dir
	projectVEnvDirName := projectDirName + "/venv"
	// create virtual environment
	cmd := exec.Command("python", "-m", "venv", projectVEnvDirName)
	err := cmd.Run()
	if err != nil {
		return false, err
	}
	return true, nil
}
