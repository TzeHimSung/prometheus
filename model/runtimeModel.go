package model

import (
	"context"
	"time"
)

type ProjectInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type FileSuffixInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ModelInfo struct {
	Id         int    `json:"id"`
	ScriptName string `json:"scriptName"`
}

type RunningModel struct {
	Id         int
	ScriptName string
	Ctx        context.Context
	CancelFunc context.CancelFunc
	LaunchTime time.Time
}