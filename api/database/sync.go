/**
 * @Description: sync database table structure
 */
package database

import (
	"prometheus/model"
	"xorm.io/xorm"
)

/**
 * @Description: sync database table structure
 * @param dbEngine: database engine
 * @return bool: sign of sync process
 * @return error: error when sync table failed
 */
func SyncTableStructure(dbEngine *xorm.Engine) (bool, error) {
	// use Sync2 but not sync
	// reason: https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-03/4.sync.html
	err := dbEngine.Sync2(new(model.Project))
	if err != nil {
		return false, err
	}
	err = dbEngine.Sync2(new(model.FileInfo))
	if err != nil {
		return false, err
	}
	err = dbEngine.Sync2(new(model.FinishedModelInfo))
	if err != nil {
		return false, err
	}
	return true, nil
}
