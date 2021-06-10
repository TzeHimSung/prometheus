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

type ProjectResultDir struct {
	ResultDirName string    `json:"resultDirName"` // project result dir name
	CreateTime    time.Time `json:"createTime"`    // project result create time
}

type PersonalEntity struct {
	PER string `json:"PER"`
	BIR string `json:"BIR"`
	LOC string `json:"LOC"`
	HJ  string `json:"HJ"`
	NAT string `json:"NAT"`
	EDU string `json:"EDU"`
	T   string `json:"T"`
	MON string `json:"MON"`
}

type EntityStruct struct {
	Entity PersonalEntity
}

type CHNNECStruct struct {
	DataList []EntityStruct
}
