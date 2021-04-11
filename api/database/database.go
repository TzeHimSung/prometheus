// Package database
// APIs related to database
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
	// DBConfigPath database configuration path
	DBConfigPath = "./config.json"
)

var (
	// dbEngine database engine
	dbEngine *xorm.Engine
)

// readDBConfig read database configuration at project root path
/**
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

// InitDatabase initial database
/**
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

// QueryProjectLog query project log from database
/**
 * @return []model.Project: project log slice
 * @return error: error
 */
func QueryProjectLog() ([]model.Project, error) {
	// get all project info from database
	projectList := make([]model.Project, 0)
	err := dbEngine.Find(&projectList)
	if err != nil {
		return nil, err
	}
	return projectList, nil
}

// AddProjectLog add project log to database
/**
 * @param projectName: project name
 * @return bool: result of adding process
 * @return error: error
 */
func AddProjectLog(projectName string) (bool, error) {
	// add project info to database
	newProject := model.Project{
		ProjectName: projectName,
	}
	_, err := dbEngine.Insert(&newProject)
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeleteProjectLog delete project log to database
/**
 * @param projectName project name
 * @return bool: result of deleting process
 * @return error: error
 */
func DeleteProjectLog(projectName string) (bool, error) {
	// delete project info from database
	project := model.Project{
		ProjectName: projectName,
	}
	_, err := dbEngine.Delete(&project)
	if err != nil {
		return false, err
	}
	return true, nil
}

// QueryProjectFileLog query project file information from database
/**
 * @param projectName: project name
 * @return []model.FileInfo: file information slice
 * @return error: error
 */
func QueryProjectFileLog(projectName string) ([]model.FileInfo, error) {
	fileList := make([]model.FileInfo, 0)
	err := dbEngine.Where("project_name = ?", projectName).Find(&fileList)
	if err != nil {
		return nil, err
	}
	return fileList, nil
}

// AddUploadFileLog add upload data log to database
/**
 * @param filename: data file name
 * @return bool: result of adding process
 * @return error: error when adding process failed
 */
func AddUploadFileLog(filename string) (bool, error) {
	// add data upload info to database
	newFile := model.FileInfo{
		FileName:    filename,
		ProjectName: model.CurrProject,
		Source:      "User upload",
	}
	_, err := dbEngine.Insert(&newFile)
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeleteUploadFileLog delete data log from database
/**
 * @param filename: data file name
 * @return bool: result of delete process
 * @return error: error when delete process failed
 */
func DeleteUploadFileLog(filename string) (bool, error) {
	// delete data upload info from database
	dataFile := model.FileInfo{
		FileName:    filename,
		ProjectName: model.CurrProject,
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
