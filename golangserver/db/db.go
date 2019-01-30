package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func ConnectDatabase() *sql.DB{
	db, err = sql.Open("mysql", "root:root@/bcg_bdpower")

	return db
}

func CloseDb() {
	db.Close()
}
