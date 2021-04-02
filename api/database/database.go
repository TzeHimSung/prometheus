/**
 * @Description: APIs related to database
 */
package database

import (
	"bufio"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/golog"
	"io"
	"os"
	"prometheus/model"
	"time"
	"xorm.io/xorm"
)

const (
	// database configuration path
	DBConfigPath = "./config.json"
)

var (
	// database engine
	dbEngine *xorm.Engine
)

/**
 * @Description: read database configuration at project root path
 * @return result: a string contains database configuration
 */
func readDBConfig() (result string) {
	// load file
	file, err := os.Open(DBConfigPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read file
	buf := bufio.NewReader(file)
	for {
		s, err := buf.ReadString('\n')
		result += s
		if err != nil {
			if err == io.EOF {
				fmt.Println("Read is ok")
				break
			} else {
				fmt.Println("ERROR:", err)
				return
			}
		}
	}

	// return database configuration
	return result
}

/**
 * @Description: initial database
 * @return bool: result of database initial process
 * @return error: error when initial database failed
 */
func InitDatabase() (bool, error) {
	golog.Info("Start database initialization progress...")

	// read database config (config.json at project root path)
	var dbConfigStr struct {
		Username string // db account username
		Password string // db account password
		Network  string // connection type
		Server   string // server location
		Port     int    // port
		Database string // db name
	}
	dbConfig := readDBConfig()
	// load database config as JSON
	err := json.Unmarshal([]byte(dbConfig), &dbConfigStr)
	if err != nil {
		return false, err
	}
	golog.Info("Database configuration loaded.")

	// generate database driver configuration
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", dbConfigStr.Username, dbConfigStr.Password, dbConfigStr.Network,
		dbConfigStr.Server, dbConfigStr.Port, dbConfigStr.Database)

	// generate database engine
	dbEngine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		return false, err
	}
	golog.Info("Database engine is generated.")

	// database engine configuration
	dbEngine.ShowSQL(true)                                      // print generated SQL in terminal
	dbEngine.SetConnMaxLifetime(100 * time.Second)              // maximum of life time of database connection
	dbEngine.SetMaxOpenConns(100)                               // maximum of open connections
	dbEngine.SetMaxIdleConns(16)                                // maximum of idle connections
	dbEngine.TZLocation, _ = time.LoadLocation("Asia/Shanghai") // set time zone

	// sync table structure
	golog.Info("Start syncing database tables...")
	_, err = SyncTableStructure(dbEngine)
	if err != nil {
		return false, err
	}
	golog.Info("Database tables are synced.")

	return true, nil
}

/**
 * @Description: query upload data log from database
 * @return []model.DataStoreInfo: upload data log slice
 * @return error: error
 */
func QueryUploadDataLog() ([]model.DataStoreInfo, error) {
	// get all data upload info from database
	dataLog := make([]model.DataStoreInfo, 0)
	err := dbEngine.Find(&dataLog)
	if err != nil {
		return nil, err
	}
	return dataLog, nil
}

/**
 * @Description: add upload data log to database
 * @param filename: data file name
 * @return bool: result of adding process
 * @return error: error when adding process failed
 */
func AddUploadDataLog(filename string) (bool, error) {
	// add data upload info to database
	newFile := model.DataStoreInfo{
		FileName: filename,
		Source:   "User upload",
		Status:   "Uploaded",
	}
	_, err := dbEngine.Insert(&newFile)
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
 * @Description: delete data log from database
 * @param filename: data file name
 * @return bool: result of delete process
 * @return error: error when delete process failed
 */
func DeleteUploadDataLog(filename string) (bool, error) {
	// delete data upload info from database
	dataFile := model.DataStoreInfo{
		FileName: filename,
	}
	_, err := dbEngine.Delete(&dataFile)
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
 * @Description: query upload model log from database
 * @return []model.ModelStoreInfo: upload model log slice
 * @return error: error
 */
func QueryUploadModelLog() ([]model.ModelStoreInfo, error) {
	// get all model upload info from database
	modelLog := make([]model.ModelStoreInfo, 0)
	err := dbEngine.Find(&modelLog)
	if err != nil {
		return nil, err
	}
	return modelLog, nil
}

/**
 * @Description: add upload model log to database
 * @param filename: model file name
 * @return bool: result of adding process
 * @return error: error when adding process failed
 */
func AddUploadModelLog(filename string) (bool, error) {
	// add model upload info to database
	newFile := model.ModelStoreInfo{
		FileName: filename,
		Source:   "User upload",
		Status:   "Uploaded",
	}
	_, err := dbEngine.Insert(&newFile)
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
 * @Description: delete model log from database
 * @param filename: model file name
 * @return bool: result of delete model log process
 * @return error: error when delete process failed
 */
func DeleteUploadModelLog(filename string) (bool, error) {
	// delete model upload info from database
	dataFile := model.ModelStoreInfo{
		FileName: filename,
	}
	_, err := dbEngine.Delete(&dataFile)
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
 * @Description: query finished model info from database
 * @return []model.FinishedModelInfo: finished model log slice
 * @return error: error
 */
func QueryFinishedModelLog() ([]model.FinishedModelInfo, error) {
	// get all finished model info from database
	modelLog := make([]model.FinishedModelInfo, 0)
	err := dbEngine.Find(&modelLog)
	if err != nil {
		return nil, err
	}
	return modelLog, nil
}

/**
 * @Description: add finished model log to database
 * @param modelID: model id
 * @param modelname: model name
 * @param launchTime: model launch time
 * @return bool: result of adding process
 * @return error: error when adding process failed
 */
func AddFinishedModelLog(modelID int, modelname string, launchTime time.Time) (bool, error) {
	// add finished model log to database
	modelLog := model.FinishedModelInfo{
		Id:         modelID,
		ScriptName: modelname,
		Status:     "Finished",
		LaunchTime: launchTime,
		FinishTime: time.Now(),
	}
	_, err := dbEngine.Insert(&modelLog)
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
 * @Description: add kill model log to database
 * @param modelID: model id
 * @param modelname: model name
 * @param launchTime: model launch time
 * @return bool: result of adding process
 * @return error: error when adding process failed
 */
func AddKilledModelLog(modelID int, modelname string, launchTime time.Time) (bool, error) {
	// add killed model log to database
	modelLog := model.FinishedModelInfo{
		Id:         modelID,
		ScriptName: modelname,
		Status:     "Killed",
		LaunchTime: launchTime,
		FinishTime: time.Now(),
	}
	_, err := dbEngine.Insert(&modelLog)
	if err != nil {
		return false, err
	}
	return true, nil
}
