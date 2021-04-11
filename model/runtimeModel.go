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
	Ctx        context.Context    // context related to model
	CancelFunc context.CancelFunc // cancel function related to model
	LaunchTime time.Time          // model launch time
}

type RunningProject struct {
	Id          int                // project id
	ProjectName string             // project name
	Ctx         context.Context    // context related to project
	CancelFunc  context.CancelFunc // cancel function related to project
	LaunchTime  time.Time          // project launch time
}
