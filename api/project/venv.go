package project

import (
	"os"
	"os/exec"
	"prometheus/api/modeltraining"
	. "prometheus/model"
	"time"
)

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
