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
	FileName   string `json:"fileName" xorm:"Varchar(255)"`
	Source     string `json:"source" xorm:"Varchar(255)"`
	Status     string `json:"status" xorm:"Varchar(255)"`
	CreateTime string `json:"createTime" xorm:"created"`
}

type ModelStoreInfo struct {
	FileName   string `json:"fileName" xorm:"Varchar(255)"`
	Source     string `json:"source" xorm:"Varchar(255)"`
	Status     string `json:"status" xorm:"Varchar(255)"`
	CreateTime string `json:"createTime" xorm:"created"`
}

type RunningModel struct {
	Id         int    `xorm:"Int"`
	ScriptName string `xorm:"Varchar(255)"`
	Ctx        *context.Context
	CancelFunc *context.CancelFunc
}
