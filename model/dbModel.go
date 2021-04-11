/**
 * @Description: data structure of database models
 */
package model

import (
	"time"
)

type Project struct {
	ProjectName string    `json:"projectName" xorm:"varchar(255)"` // project name
	CreateTime  time.Time `json:"createTime" xorm:"created"`       // project create time
}

type FileInfo struct {
	ProjectName string    `json:"projectName" xorm:"varchar(255)"` // project name
	FileName    string    `json:"fileName" xorm:"varchar(255)"`    // file name
	Source      string    `json:"source" xorm:"varchar(255)"`      // upload source
	CreateTime  time.Time `json:"createTime" xorm:"created"`       // file create time
}

type FinishedModelInfo struct {
	Id         int       `json:"id" xorm:"int"`                  // model id
	ScriptName string    `json:"scriptName" xorm:"varchar(255)"` // model name
	Status     string    `json:"status" xorm:"varchar(255)"`     // status
	LaunchTime time.Time `json:"launchTime" xorm:"created"`      // model launch time
	FinishTime time.Time `json:"finishTime" xorm:"created"`      // model finish time
}

type FinishedProjectInfo struct {
	Id          int       `json:"id" xorm:"int"`                   // project id
	ProjectName string    `json:"projectName" xorm:"varchar(255)"` // project name
	Status      string    `json:"status" xorm:"varchar(255)"`      // project status
	LaunchTime  time.Time `json:"launchTime" xorm:"DateTime"`      // project launch time
	FinishTime  time.Time `json:"finishTime" xorm:"created"`       // project finish time
}
