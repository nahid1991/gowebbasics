package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func ConnectDatabase() {
	fmt.Println("Database connected.")
	db, err = sql.Open("mysql", "root:root@/bcg_bdpower")
}
