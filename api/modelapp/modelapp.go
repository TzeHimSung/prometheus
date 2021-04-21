/**
 * @Description: APIs related to modelapp page
 */
package modelapp

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"prometheus/api/project"
	"prometheus/model"
)

// GetProjectResultDir get project result dir list
/**
 * @return []string: project result dir list
 */
func GetProjectResultDir() ([]string, error) {
	dirList := make([]string, 0)
	files, err := ioutil.ReadDir(model.ModelOutputPath)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			dirList = append(dirList, file.Name())
		}
	}
	return dirList, nil
}

// LoadProjectResult
/**
 * @param dirName
 * @return string
 */
func LoadProjectResult(dirName string) (string, error) {
	// read output content
	file, err := os.Open(filepath.Join(project.OutputRootPath, dirName, "output.txt"))
	if err != nil {
		return "", err
	}
	defer file.Close()
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(fileContent), nil
}
