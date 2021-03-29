package model

import (
	"time"
)

type DataStoreInfo struct {
	FileName   string    `json:"fileName" xorm:"varchar(255)"`
	Source     string    `json:"source" xorm:"varchar(255)"`
	Status     string    `json:"status" xorm:"varchar(255)"`
	CreateTime time.Time `json:"createTime" xorm:"created"`
}

type ModelStoreInfo struct {
	FileName   string    `json:"fileName" xorm:"varchar(255)"`
	Source     string    `json:"source" xorm:"varchar(255)"`
	Status     string    `json:"status" xorm:"varchar(255)"`
	CreateTime time.Time `json:"createTime" xorm:"created"`
}

type RunningModelInfo struct {
	Id         int       `json:"id" xorm:"int"`
	ScriptName string    `json:"scriptName" xorm:"varchar(255)"`
	Status     string    `json:"status" xorm:varchar(255)`
	LaunchTime time.Time `json:"launchTime" xorm:"created"`
	FinishTime time.Time `json:"finishTime" xorm:"created"`
}
