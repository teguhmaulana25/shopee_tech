package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var Db *sql.DB
var err error
var UpdateResult DBUpdate

type DBUpdate struct {
	ID       int   `json:"id"`
	Affected int64 `json:"affected"`
}

func DbConnect() {
	var dbUser = os.Getenv("DBUSER")
	var dbPass = os.Getenv("DBPASS")
	var dbHost = os.Getenv("DBHOST")
	var dbDatabase = os.Getenv("DBDATABASE")

	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&collation=utf8mb4_unicode_ci", dbUser, dbPass, dbHost, dbDatabase)
	Db, err = sql.Open("mysql", dbConn)
	if err != nil {
		log.Panicf("Error: %s", err)
	}

	err = Db.Ping()
	if err != nil {
		log.Panicf("Error: %s", err)
	}
	log.Info("Database connected.")

}

func Check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
