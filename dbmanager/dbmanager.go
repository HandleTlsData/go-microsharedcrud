package dbmanager

import (
	"database/sql"
	"fmt"
)

var hostname string
var username string
var password string
var dbName string

var DBCon *sql.DB

func connect() {
	var err error
	//username:password@tcp(host:port)/dbname
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
	fmt.Println("Establishing connection " + connStr)
	DBCon, err = sql.Open("mysql", connStr)

	if err != nil {
		panic(err.Error())
	}

}
