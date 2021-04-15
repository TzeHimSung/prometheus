// Package model
/**
 * @Description: global const or variable
 */
package model

import "path/filepath"

const (
	// TimeFormat standard time format
	TimeFormat = "2006-01-02 15:04:05"
)

var (
	// CurrProject current project
	CurrProject = "Sample Project"
	// ModelID model id counter
	ModelID = 0
	// RunningModelList running model list
	RunningModelList = make([]RunningModel, 0)
	// ProjectPath project file path
	ProjectPath = filepath.Join("uploads", "project")
	// ModelOutputPath model output path
	ModelOutputPath = filepath.Join(".", "runmodel", "output")
)
