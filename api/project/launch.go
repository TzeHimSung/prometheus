// Package project
/**
 * @Description: APIs related to project
 */
package project

import (
	"github.com/kataras/golog"
	"os"
	"os/exec"
	"path/filepath"
	"prometheus/api/database"
	. "prometheus/model"
	"strconv"
	"strings"
	"time"
)

const (
	// OutputTimeFormat model file output path time format
	OutputTimeFormat = "2006-01-02-15-04-05"
)

var (
	// ProjectID project id counter
	ProjectID = 0
	// RunningProjectList running project list
	RunningProjectList = make([]RunningProject, 0)
	// OutputRootPath model file output root path
	OutputRootPath = filepath.Join("runmodel", "output")
)

// LaunchProject launch specific project
/**
 * @param projectName: project name
 * @param projectID: project id
 * @param ctx: context related to project
 */
//func LaunchProject(projectName string, projectID int, quitChan chan int) {
//	for {
//		select {
//		case <-quitChan:
//			// quit launch goroutine
//			golog.Warn("Project " + projectName + " has been canceled.")
//			return
//		default:
//			golog.Info("Launching project: " + projectName)
//
//			// create output dir
//			golog.Info("Create output dir...")
//			outputPath := OutputRootPath + "/proj-" +
//				strings.Replace(projectName, ".", "-", -1) +
//				"-" + time.Now().Format(OutputTimeFormat)
//			_, err := os.Stat(outputPath)
//			if err != nil {
//				err = os.Mkdir(outputPath, 0666)
//				if err != nil {
//					golog.Error("Can not create dir: " + outputPath + ", please check dir name.")
//					return
//				}
//			}
//			golog.Info("Output path check passed. Start project...")
//
//			// launch project
//			// need to check main.py exists or not
//			launchFilePath := filepath.Join(ProjectPath, projectName, "main.py")
//			venvActiveCmd := filepath.Join(ProjectPath, projectName, "venv", "Scripts", "activate.bat")
//			cmd := exec.Command(venvActiveCmd, "&&", "python", launchFilePath)
//
//			err = cmd.Start()
//			if err != nil {
//				golog.Error("Error in project " + projectName + " launching process")
//				panic(err)
//			}
//			golog.Info("Project is running, pid: ", cmd.Process.Pid)
//
//			// add running project record
//			RunningProjectList = append(RunningProjectList, RunningProject{
//				Id:          ProjectID,
//				Pid:         cmd.Process.Pid,
//				ProjectName: projectName,
//				LaunchTime:  time.Now(),
//				QuitChan:    quitChan,
//			})
//
//			// wait for process end
//			err = cmd.Wait()
//			if err != nil {
//				golog.Error(err)
//			}
//
//			// remove running project log
//			var projectIdx = 0
//			for idx, runningProject := range RunningProjectList {
//				if runningProject.Id == projectID {
//					// add database log
//					_, err := database.AddFinishedProjectLog(
//						runningProject.Id,
//						runningProject.Pid,
//						runningProject.ProjectName,
//						runningProject.LaunchTime)
//					if err != nil {
//						panic(err)
//					}
//					projectIdx = idx
//					break
//				}
//			}
//			RunningProjectList = append(RunningProjectList[:projectIdx],
//				RunningProjectList[projectIdx+1:]...)
//
//			return
//		}
//	}
//}

// LaunchProject another version of launch project
// redirect stdout to output.txt
func LaunchProject(projectName string, projectID int, quitChan chan int) {
	for {
		select {
		case <-quitChan:
			// quit launch goroutine
			golog.Warn("Project " + projectName + " has been canceled.")
			return
		default:
			golog.Info("Launching project: " + projectName)

			// create output dir
			golog.Info("Create output dir...")
			outputFolderName := "proj-" +
				strings.Replace(projectName, ".", "-", -1) +
				"-" + time.Now().Format(OutputTimeFormat)
			outputPath := filepath.Join(OutputRootPath, outputFolderName)
			_, err := os.Stat(outputPath)
			if err != nil {
				err = os.Mkdir(outputPath, 0666)
				if err != nil {
					golog.Error("Can not create dir: " + outputPath + ", please check dir name.")
					return
				}
			}
			golog.Info("Output path check passed. Start project...")

			// launch project
			// need to check main.py exists or not
			launchFilePath := filepath.Join(ProjectPath, projectName, "main.py")
			venvActiveCmd := filepath.Join(ProjectPath, projectName, "venv", "Scripts", "activate.bat")
			cmd := exec.Command(venvActiveCmd, "&&", "python", launchFilePath, outputFolderName)

			// add running project record
			RunningProjectList = append(RunningProjectList, RunningProject{
				Id:          ProjectID,
				Pid:         0,
				ProjectName: projectName,
				LaunchTime:  time.Now(),
				QuitChan:    quitChan,
			})

			out, err := cmd.Output()
			if err != nil {
				golog.Error("Error in project " + projectName + " launching process")
				panic(err)
			}

			// create output file
			f, err := os.Create(outputPath + "/output.txt")
			defer f.Close()
			if err != nil {
				panic(err)
			} else {
				// write output
				_, err := f.Write(out)
				if err != nil {
					panic(err)
				}
				golog.Info("Output file is created.")
			}

			// remove running project log
			var projectIdx = 0
			for idx, runningProject := range RunningProjectList {
				if runningProject.Id == projectID {
					// add database log
					_, err := database.AddFinishedProjectLog(
						runningProject.Id,
						runningProject.Pid,
						runningProject.ProjectName,
						runningProject.LaunchTime)
					if err != nil {
						panic(err)
					}
					projectIdx = idx
					break
				}
			}
			RunningProjectList = append(RunningProjectList[:projectIdx],
				RunningProjectList[projectIdx+1:]...)

			return
		}
	}
}

// KillProcessWithPid kill process using taskkill
/**
 * @param pid: process id
 * @return error: error
 */
func KillProcessWithPid(pid int) error {
	cmd := exec.Command("taskkill", "/PID", strconv.Itoa(pid), "/F")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
