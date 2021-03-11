package modelstore

import (
	. "prometheus/datastructure"
)

/**
 * @description: get stored model information
 * @param none
 * @return []ModelStoreInfo
 */
func GetModelStoreInfo() []ModelStoreInfo {
	return []ModelStoreInfo{
		{
			FileName:   "modelName1",
			Source:     "user upload",
			Status:     "上传中",
			CreateTime: "2021-03-08 00:00:00",
		},
		{
			FileName:   "modelName2",
			Source:     "user upload",
			Status:     "已上传",
			CreateTime: "2021-03-08 00:00:00",
		},
		{
			FileName:   "modelName3",
			Source:     "user upload",
			Status:     "上传中",
			CreateTime: "2021-03-08 00:00:00",
		},
		{
			FileName:   "modelName4",
			Source:     "user upload",
			Status:     "已上传",
			CreateTime: "2021-03-08 00:00:00",
		},
		{
			FileName:   "modelName5",
			Source:     "user upload",
			Status:     "上传中",
			CreateTime: "2021-03-08 00:00:00",
		},
	}
}
