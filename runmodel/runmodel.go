package runmodel

import (
	"github.com/kataras/golog"
	"os"
	"os/exec"
)

const (
	OutputRootPath = "./runmodel/output/"
)

func Launch(filename string) {
	golog.Info("Launch file: " + filename)
	filepath := "./uploads/model/" + filename
	cmd := exec.Command("python", filepath)
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	outputPath := OutputRootPath + "test"
	_, err = os.Stat(outputPath)
	if err != nil {
		err = os.Mkdir(outputPath, 0666)
	}
	f, err := os.Create(outputPath + "/output.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	} else {
		_, err := f.Write(out)
		if err != nil {
			panic(err)
		}
	}
}
