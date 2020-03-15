package repo

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	mySqlDB *sql.DB
)

const (
	mySqlDrive            = "mysql"
	mySqlConnectionString = "root:19881220@/vegetablesmarket?charset=utf8"
)

func init() {
	db, err := sql.Open(mySqlDrive, mySqlConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	mySqlDB = db
}
