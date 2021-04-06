/**
 * @Description: APIs related to project
 */
package project

import (
	"github.com/kataras/golog"
	"os"
	"os/exec"
	"prometheus/api/modeltraining"
	. "prometheus/model"
	"time"
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

/**
 * @Description: install virtual environment required package
 * @param projectName: project name
 * @return bool: sign of installation process
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
		"pip", "install", "-r", "uploads/project/test/requirements.txt")
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

/**
 * @Description: get pip list info
 * @param projectName: project name
 * @return bool: sign of query process
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
