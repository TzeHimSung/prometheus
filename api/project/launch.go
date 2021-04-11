// Package project
/**
 * @Description: APIs related to project
 */
package project

import (
	"context"
	"github.com/kataras/golog"
	"os"
	"os/exec"
	"prometheus/api/database"
	"prometheus/model"
	"strings"
	"time"
)

const (
	// OutputRootPath model file output root path
	OutputRootPath = "./runmodel/output/"
	// OutputTimeFormat model file output path time format
	OutputTimeFormat = "2006-01-02-15-04-05"
)

func LaunchProject(projectName string, projectID int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			golog.Warn("Project " + projectName + " is canceled.")
			return
		default:
			golog.Info("Launch project: " + projectName)

			// create output dir
			outputPath := OutputRootPath + "proj-" +
				strings.Replace(projectName, ".", "-", -1) +
				"-" + time.Now().Format(OutputTimeFormat)
			_, err := os.Stat(outputPath)
			if err != nil {
				err = os.Mkdir(outputPath, 0666)
				if err != nil {
					golog.Error("Can not create dir: " + outputPath + ", please check dir name.")
					return
				}
			}
			golog.Info("Output path check passed. Launching project...")

			// launch project
			// need to check main.py exists or not
			launchFilePath := model.ProjectPath + "/" + projectName + "/main.py"
			venvActiveCmd := "uploads\\project\\" + projectName + "\\venv\\Scripts\\activate.bat"
			cmd := exec.Command(venvActiveCmd, "&&", "python", launchFilePath)
			golog.Info("Project is launching")
			out, err := cmd.Output()
			if err != nil {
				golog.Error("Error in project " + projectName + " launching process")
				panic(err)
			}
			golog.Info("Project launch process finished. Create output file...")

			// create output file
			f, err := os.Create(outputPath + "/output.txt")
			defer f.Close()
			if err != nil {
				panic(err)
			} else {
				// save output to output.txt
				_, err := f.Write(out)
				if err != nil {
					panic(err)
				}
				golog.Info("Output file is created.")
			}

			// remove running project log
			var projectIdx int = 0
			for idx, runningProject := range model.RunningProjectList {
				if runningProject.Id == projectID {
					_, err := database.AddFinishedProjectLog(runningProject.Id,
						runningProject.ProjectName,
						runningProject.LaunchTime)
					if err != nil {
						panic(err)
					}
					projectIdx = idx
					break
				}
			}
			model.RunningProjectList = append(model.RunningProjectList[:projectIdx],
				model.RunningProjectList[projectIdx+1:]...)

			// initiative return is needed, or it will run continuously
			return
		}
	}
}
