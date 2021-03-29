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
	DBConfigPath = "./config.json"
)

var dbEngine *xorm.Engine

func readDBConfig() (result string) {
	file, err := os.Open(DBConfigPath)
	defer file.Close()
	if err != nil {
		panic(err)
	}
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
	return result
}

func InitDatabase() (bool, error) {
	golog.Info("Start database initialization progress...")
	// read database config (config.json at project root path)
	var dbConfigStr struct {
		Username string
		Password string
		Network  string
		Server   string
		Port     int
		Database string
	}
	dbConfig := readDBConfig()
	err := json.Unmarshal([]byte(dbConfig), &dbConfigStr)
	if err != nil {
		return false, err
	}
	golog.Info("Database configuration loaded.")

	// generate database driver configuration
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", dbConfigStr.Username, dbConfigStr.Password, dbConfigStr.Network,
		dbConfigStr.Server, dbConfigStr.Port, dbConfigStr.Database)

	// get database engine
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

	// test create table
	golog.Info("Start syncing database tables...")
	err = dbEngine.Sync2(new(model.DataStoreInfo))
	if err != nil {
		return false, err
	}
	err = dbEngine.Sync2(new(model.ModelStoreInfo))
	if err != nil {
		return false, err
	}
	err = dbEngine.Sync2(new(model.RunningModelInfo))
	if err != nil {
		return false, err
	}
	golog.Info("Database tables are synced.")

	return true, nil
}

func QueryUploadDataLog() ([]model.DataStoreInfo, error) {
	// get all data upload info from database
	dataLog := make([]model.DataStoreInfo, 0)
	err := dbEngine.Find(&dataLog)
	if err != nil {
		return nil, err
	}
	return dataLog, nil
}

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

func QueryUploadModelLog() ([]model.ModelStoreInfo, error) {
	// get all model upload info from database
	modelLog := make([]model.ModelStoreInfo, 0)
	err := dbEngine.Find(&modelLog)
	if err != nil {
		return nil, err
	}
	return modelLog, nil
}

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

func QueryFinishedModelLog() ([]model.RunningModelInfo, error) {
	// get all finished model info from database
	modelLog := make([]model.RunningModelInfo, 0)
	err := dbEngine.Find(&modelLog)
	if err != nil {
		return nil, err
	}
	return modelLog, nil
}

func AddFinishedModelLog(modelID int, modelname string, launchTime time.Time) (bool, error) {
	// add finished model log to database
	modelLog := model.RunningModelInfo{
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

func AddKilledModelLog(modelID int, modelname string, launchTime time.Time) (bool, error) {
	// add killed model log to database
	modelLog := model.RunningModelInfo{
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
