/**
 * @Description: APIs related to modeltraining page
 */
package modeltraining

import (
	"context"
	"github.com/kataras/golog"
	"os"
	"os/exec"
	"prometheus/api/database"
	. "prometheus/model"
	"strings"
	"time"
)

const (
	// model file output root path
	OutputRootPath = "./runmodel/output/"
	// model file output path time format
	OutputTimeFormat = "2006-01-02-15-04-05"
)

/**
 * @Description: launch specific model
 * @param filename: model file name
 * @param modelID: model id
 * @param ctx: context
 */
func LaunchModel(ctx context.Context, filename string, modelID int) {
	for {
		select {
		// return goroutine when received context done single
		case <-ctx.Done():
			golog.Warn("Model " + filename + "is canceled.")
			return
		default:
			golog.Info("Launch Model file: " + filename)

			// create output dir
			outputPath := OutputRootPath + "proj-" + strings.Replace(filename, ".", "-", -1) + "-" + time.Now().Format(OutputTimeFormat)
			_, err := os.Stat(outputPath)
			if err != nil {
				err = os.Mkdir(outputPath, 0666)
				if err != nil {
					golog.Error("Can not create dir: " + outputPath + ", please check dir name.")
					return
				}
			}
			golog.Info("Output path check passed. Launching model...")

			// launch model
			modelPath := "./uploads/model/" + filename
			cmd := exec.Command("python", modelPath)
			golog.Info("Model is running...")
			out, err := cmd.Output()
			if err != nil {
				golog.Info("Error in model launching process.")
				panic(err)
			}
			golog.Info("Model launched. Create output file...")

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

			// remove running model log
			var modelIdx int = 0
			for idx, runningModel := range RunningModelList {
				if runningModel.Id == modelID {
					// add finished log
					_, err := database.AddFinishedModelLog(runningModel.Id, runningModel.ScriptName, runningModel.LaunchTime)
					if err != nil {
						panic(err)
					}
					modelIdx = idx
					break
				}
			}
			RunningModelList = append(RunningModelList[:modelIdx],
				RunningModelList[modelIdx+1:]...)

			// initiative return, or it will run continuously
			return
		}
	}
}
