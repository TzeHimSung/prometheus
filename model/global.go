// Package model
/**
 * @Description: global const or variable
 */
package model

const (
	// TimeFormat standard time format
	TimeFormat = "2006-01-02 15:04:05"
	// ProjectPath project file path
	ProjectPath = "./uploads/project"
	// ModelOutputPath model output path
	ModelOutputPath = "./runmodel/output"
)

var (
	// CurrProject current project
	CurrProject = "Sample Project"
	// ModelID model id counter
	ModelID = 0
	// ProjectID project id counter
	ProjectID = 0
	// RunningModelList running model list
	RunningModelList = make([]RunningModel, 0)
	// RunningProjectList running project list
	RunningProjectList = make([]RunningProject, 0)
)
