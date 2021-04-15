/**
 * @Description: data structure of runtime models
 */
package model

import (
	"context"
	"time"
)

type ModelInfo struct {
	Id         int    `json:"id"`         // model id
	ScriptName string `json:"scriptName"` // model file name
}

type RunningModel struct {
	Id         int                // model id
	ScriptName string             // model name
	CancelFunc context.CancelFunc // cancel function related to model
	LaunchTime time.Time          // model launch time
}

type RunningProject struct {
	Id          int       // project id
	Pid         int       // process id of project
	ProjectName string    // project name
	LaunchTime  time.Time // project launch time
	QuitChan    chan int  // project quit channel
}
