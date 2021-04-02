/**
 * @Description: data structure of runtime models
 */
package model

import (
	"context"
	"time"
)

type ProjectInfo struct {
	Id   int    `json:"id"`   // project id
	Name string `json:"name"` // project name
}

type FileSuffixInfo struct {
	Id   int    `json:"id"`   // file suffix id
	Name string `json:"name"` // file suffix name
}

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
