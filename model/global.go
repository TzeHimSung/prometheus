/**
 * @Description: global const or variable
 */
package model

const (
	// standard time format
	TimeFormat = "2006-01-02 15:04:05"
	// data file path
	DataPath = "./uploads/data"
	// model file path
	ModelPath = "./uploads/model"
)

var (
	// model id counter
	ModelID = 0
	// running model list
	RunningModelList = make([]RunningModel, 0)
)
