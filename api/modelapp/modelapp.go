/**
 * @Description: APIs related to modelapp page
 */
package modelapp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"prometheus/api/project"
	"prometheus/model"
	"runtime"
	"syscall"
	"time"
)

// GetProjectResultDir get project result dir list
/**
 * @return []string: project result dir list
 */
func GetProjectResultDir() ([]model.ProjectResultDir, error) {
	dirList := make([]model.ProjectResultDir, 0)
	files, err := ioutil.ReadDir(model.ModelOutputPath)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			var fileCreateTime time.Time
			// get dir create time
			osType := runtime.GOOS
			dirInfo, _ := os.Stat(filepath.Join(project.OutputRootPath, file.Name()))
			if osType == "windows" {
				wFileSys := dirInfo.Sys().(*syscall.Win32FileAttributeData)
				tNanSeconds := wFileSys.CreationTime.Nanoseconds()
				// convert nano second to second
				tSec := tNanSeconds / 1e9
				fileCreateTime = time.Unix(tSec, 0)
			} else {
				fileCreateTime = time.Now()
			}
			// add to dir list
			dirList = append(dirList, model.ProjectResultDir{
				ResultDirName: file.Name(),
				CreateTime:    fileCreateTime,
			})
		}
	}
	return dirList, nil
}

// LoadProjectResult
/**
 * @param dirName
 * @return string
 */
func LoadProjectResult(dirName string) (model.CHNNECStruct, error) {
	var resultStruct model.CHNNECStruct
	// read output content
	file, err := os.Open(filepath.Join(project.OutputRootPath, dirName, "output.json"))
	if err != nil {
		return resultStruct, err
	}
	defer file.Close()

	// read file
	var result string
	buf := bufio.NewReader(file)
	for {
		s, err := buf.ReadString('\n')
		result += s
		if err != nil {
			if err == io.EOF {
				fmt.Println("Read is ok")
				break
			} else {
				fmt.Println("ERROR:", err)
				return resultStruct, err
			}
		}
	}
	err = json.Unmarshal([]byte(result), &resultStruct)
	if err != nil {
		return resultStruct, err
	}

	return resultStruct, nil
}
