package dboperation

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"time"
)

const (
	DBConfigPath = "./config.json"
)

type User struct {
	Id   sql.NullInt64  `db:"id"`
	Name sql.NullString `db:"name"`
}

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

func DbTest() {
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
		panic(err)
	}
	fmt.Println(dbConfigStr)
	user := new(User)
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", dbConfigStr.Username, dbConfigStr.Password, dbConfigStr.Network,
		dbConfigStr.Server, dbConfigStr.Port, dbConfigStr.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(100 * time.Second)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	res := db.QueryRow("SELECT * FROM table1")
	if err := res.Scan(&user.Id, &user.Name); err != nil {
		panic(err)
	}
	fmt.Println(*user)
}
