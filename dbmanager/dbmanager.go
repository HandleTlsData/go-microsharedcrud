package dbmanager

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBEntity struct {
	ID          int    `json:"ID" db:"id"`
	Name        string `json:"Name" db:"Name"`
	Description string `json:"Description" db:"Description"`
	Status      string `json:"Status" db:"Status"`
	Uid         int    `json:"uid" db:"uid"`
}

type DBConfig struct {
	Hostname string
	Username string
	Password string
	DbName   string
	DBCon    *sqlx.DB
}

var AlphaDBConfig DBConfig
var BetaDBConfig DBConfig

func DeleteEntity(z *DBConfig, entID int) error {
	var err error
	entity, err := GetEntityByID(z, int64(entID))

	if err == nil {
		result, err := z.DBCon.NamedExec("DELETE FROM pendingData WHERE id=:first",
			map[string]interface{}{"first": entity.ID})
		if err != nil {
			log.Println(err.Error())
			return err
		}
		log.Println(result)

		return nil
	}
	return err
}

func StoreEntityBeta(z *DBConfig, newEntity DBEntity) error {
	var err error

	entity, err := GetEntityByID(z, int64(newEntity.ID))
	if err != nil {
		if err.Error() == "no records found" {
			result, err := z.DBCon.NamedExec("INSERT INTO pendingData (`id`, `Name`, `Description`, `Status`) VALUES (:zero, :first, :second, :third)",
				map[string]interface{}{"zero": newEntity.ID, "first": newEntity.Name, "second": newEntity.Description, "third": newEntity.Status})
			if err != nil {
				log.Println(err.Error())
				return err
			}

			log.Println(result)

			return nil

		} else {
			return err
		}
	} else {
		if entity.Description == newEntity.Description && entity.Status == newEntity.Status {
			return nil
		} else {
			result, err := z.DBCon.NamedExec("UPDATE pendingData SET Description=:first, Status=:second WHERE id=:third",
				map[string]interface{}{"first": newEntity.Description, "second": newEntity.Status, "third": newEntity.ID})
			if err != nil {
				log.Println(err.Error())
				return err
			}
			log.Println(result)
			return nil
		}
	}

}

func StoreEntityAlpha(z *DBConfig, newEntity DBEntity) error {
	var err error
	// rows, err := z.DBCon.NamedQuery("SELECT * FROM pendingData WHERE id=:first", map[string]interface{}{"first": strconv.Itoa(newEntity.ID)})
	entity, err := GetEntityByID(z, int64(newEntity.ID))

	if err != nil {
		if err.Error() == "no records found" {
			result, err := z.DBCon.NamedExec("INSERT INTO pendingData (`id`, `Name`, `Description`, `Status`) VALUES (:zero, :first, :second, :third)",
				map[string]interface{}{"zero": newEntity.ID, "first": newEntity.Name, "second": newEntity.Description, "third": newEntity.Status})
			if err != nil {
				log.Println(err.Error())
				return err
			}

			log.Println(result)

			return nil

		} else {
			return err
		}
	} else {
		if entity.Name == newEntity.Name && entity.Status == newEntity.Status {
			return nil
		} else {
			result, err := z.DBCon.NamedExec("UPDATE pendingData SET Name=:first, Status=:second WHERE id=:third",
				map[string]interface{}{"first": newEntity.Name, "second": newEntity.Status, "third": newEntity.ID})
			if err != nil {
				log.Println(err.Error())
				return err
			}
			log.Println(result)
			return nil
		}
	}

}

func GetEntityByID(z *DBConfig, rID int64) (DBEntity, error) {
	var err error
	entity := DBEntity{}

	err = verifyConnection(z)
	if err != nil {
		fmt.Println(err.Error())
		return entity, err
	}

	rows, err := z.DBCon.NamedQuery("SELECT * FROM pendingData WHERE id=:first", map[string]interface{}{"first": rID})
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
