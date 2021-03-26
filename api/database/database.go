package database

import (
	"bufio"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"prometheus/model"
	"time"
	"xorm.io/xorm"
)

const (
	DBConfigPath = "./config.json"
	TimeFormat   = "2006-01-02 15:04:05"
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

	// generate database driver configuration
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", dbConfigStr.Username, dbConfigStr.Password, dbConfigStr.Network,
		dbConfigStr.Server, dbConfigStr.Port, dbConfigStr.Database)

	// get database engine
	dbEngine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		return false, err
	}

	// database engine configuration
	dbEngine.ShowSQL(true)                                      // print generated SQL in terminal
	dbEngine.SetConnMaxLifetime(100 * time.Second)              // maximum of life time of database connection
	dbEngine.SetMaxOpenConns(100)                               // maximum of open connections
	dbEngine.SetMaxIdleConns(16)                                // maximum of idle connections
	dbEngine.TZLocation, _ = time.LoadLocation("Asia/Shanghai") // set time zone

	// test create table
	err = dbEngine.Sync2(new(model.DataStoreInfo))
	if err != nil {
		return false, err
	}
	err = dbEngine.Sync2(new(model.ModelStoreInfo))
	if err != nil {
		return false, err
	}
	err = dbEngine.Sync2(new(model.RunningModel))
	if err != nil {
		return false, err
	}
	return true, nil
}

func AddUploadDataLog(filename string) (bool, error) {
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
	dataFile := model.DataStoreInfo{
		FileName: filename,
	}
	_, err := dbEngine.Delete(&dataFile)
	if err != nil {
		return false, err
	}
	return true, nil
}

func AddUploadModelLog(filename string) (bool, error) {
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
	dataFile := model.ModelStoreInfo{
		FileName: filename,
	}
	_, err := dbEngine.Delete(&dataFile)
	if err != nil {
		return false, err
	}
	return true, nil
}
