package dbmanager

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBEntity struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"Name" db:"Name"`
	Description string `json:"Description" db:"Description"`
	Status      string `json:"Status" db:"Status"`
}

var hostname string = "95.214.55.115"
var username string = "testdb"
var password string = "2pRSTXAMBh5wLMtF"
var dbName string = "testdb"

var DBCon *sqlx.DB

func Connect() {
	var err error
	//username:password@tcp(host:port)/dbname
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
	fmt.Println("Establishing connection " + connStr)
	DBCon, err = sqlx.Open("mysql", connStr)

	if err != nil {
		fmt.Println(err.Error())
	}

	if err = DBCon.Ping(); err != nil {
		DBCon.Close()
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Connected")

	entity := []DBEntity{}
	err = DBCon.Select(&entity, "SELECT * FROM pendingData LIMIT 1")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", entity)

}
