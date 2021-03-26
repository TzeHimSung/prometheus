package model

import "context"

type ProjectInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type FileSuffixInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DataStoreInfo struct {
	FileName   string `json:"fileName" xorm:"varchar(255)"`
	Source     string `json:"source" xorm:"varchar(255)"`
	Status     string `json:"status" xorm:"varchar(255)"`
	CreateTime string `json:"createTime" xorm:"created"`
}

type ModelStoreInfo struct {
	FileName   string `json:"fileName" xorm:"varchar(255)"`
	Source     string `json:"source" xorm:"varchar(255)"`
	Status     string `json:"status" xorm:"varchar(255)"`
	CreateTime string `json:"createTime" xorm:"created"`
}

type RunningModel struct {
	Id         int    `xorm:"Int"`
	ScriptName string `xorm:"varchar(255)"`
	Ctx        *context.Context
	CancelFunc *context.CancelFunc
}
