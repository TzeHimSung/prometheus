package datastructure

type ProjectInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type FileSuffixInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DataStoreInfo struct {
	FileName   string `json:"fileName"`
	Source     string `json:"source"`
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
}

type ModelStoreInfo struct {
	FileName   string `json:"fileName"`
	Source     string `json:"source"`
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
}
