package dbmanager

import (
	"errors"
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

type DBConfig struct {
	Hostname string
	Username string
	Password string
	DbName   string
	DBCon    *sqlx.DB
}

func GetEntityByName(z *DBConfig, rName string) (DBEntity, error) {
	var err error
	entity := DBEntity{}

	err = verifyConnection(z)
	if err != nil {
		fmt.Println(err.Error())
		return entity, err
	}

	rows, err := z.DBCon.NamedQuery("SELECT * FROM pendingData WHERE Name=:first", map[string]interface{}{"first": rName})
	if err != nil {
		fmt.Println(err.Error())
		return entity, err
	}

	for rows.Next() {
		err = rows.StructScan(&entity)
		if err != nil {
			fmt.Println(err)
			return entity, err
		}
		fmt.Printf("%+v\n", entity)
		return entity, nil
	}

	return entity, errors.New("no records found")
}

func verifyConnection(z *DBConfig) error {
	var err error
	if z.DBCon != nil {
		if err = z.DBCon.Ping(); err != nil {
			Disconnect(z)
			fmt.Println("Re-establishing connection")
			err = Connect(z)
			return err
		} else {
			fmt.Println("Still connected")
			return nil
		}
	} else {
		err = Connect(z)
		return err
	}
}

func Disconnect(z *DBConfig) {
	if z.DBCon != nil {
		z.DBCon.Close()
	}
	fmt.Println("Connection closed")

}

func Connect(z *DBConfig) error {
	var err error
	//username:password@tcp(host:port)/dbname
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", z.Username, z.Password, z.Hostname, z.DbName)
	fmt.Println("Establishing connection " + connStr)
	z.DBCon, err = sqlx.Open("mysql", connStr)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if err = z.DBCon.Ping(); err != nil {
		z.DBCon.Close()
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Connected")
	return nil
}

var CurrentDBConfig DBConfig
var AlphaDBConfig DBConfig
var BetaDBConfig DBConfig
