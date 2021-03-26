package modeltraining

import (
	"context"
	"github.com/kataras/golog"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	//DataRootPath     = "./uploads/data/"
	OutputRootPath   = "./runmodel/output/"
	OutputTimeFormat = "2006-01-02-15-04-05"
)

func LaunchModel(filename string) {
	golog.Info("Launch Model file: " + filename)

	// create output path
	// path name should be replaced
	outputPath := OutputRootPath + "model-" + strings.Replace(filename, ".", "-", -1) + "-" + time.Now().Format(OutputTimeFormat)
	_, err := os.Stat(outputPath)
	if err != nil {
		err = os.Mkdir(outputPath, 0666)
	}
	golog.Info("Output path check passed. Launching model...")

	// launch model
	modelPath := "./uploads/model/" + filename
	cmd := exec.Command("python", modelPath)
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
}

func LaunchTest(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			golog.Println("Launch process is canceled.")
			return
		default:
			golog.Println("Launch process is running")
			time.Sleep(1 * time.Second)
		}
	}
}
