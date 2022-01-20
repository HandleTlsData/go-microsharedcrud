package gdbmanager

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"sharedcrud/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var SqlDB *sql.DB

const StrNoRecords = "record not found"

func InitDB() {
	var err error
	dsn := fmt.Sprintf("testdb:eDwM8KEJpGWX2LRt@tcp(149.34.51.91:3306)/testdb?charset=utf8mb4&parseTime=True")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
	}
	sqldb, err := db.DB()
	sqldb.SetMaxIdleConns(5)
	sqldb.SetMaxOpenConns(30)
	sqldb.SetConnMaxLifetime(time.Hour)
}

func GetEntityByName(userName string) (models.Entity, error) {
	var user models.Entity
	var count int64
	err := db.Model(&models.Entity{}).Where(models.Entity{Name: userName}).First(&user).Count(&count).Error

	if err != nil {
		return models.Entity{}, err
	}

	if count > 0 {
		log.Printf("%+v\n", user)
		return user, nil
	}

	return models.Entity{}, fmt.Errorf(StrNoRecords)
}

func GetEntityByID(userID int64) (models.Entity, error) {
	var user models.Entity
	var count int64
	db.Model(&models.Entity{}).Where(models.Entity{ID: userID}).First(&user).Count(&count)

	if count > 0 {
		log.Printf("%+v\n", user)
		return user, nil
	}

	return models.Entity{}, fmt.Errorf(StrNoRecords)
}

func DeleteEntityByName(userName string) error {
	return db.Delete(&models.Entity{}, models.Entity{Name: userName}).Error
}

func StoreEntity(newEntity models.Entity) error {
	var user models.Entity
	var count int64
	err := db.Model(&models.Entity{}).Where(models.Entity{Name: newEntity.Name}).First(&user).Count(&count).Error

	if err.Error() == StrNoRecords {
		err := db.Model(&models.Entity{}).Create(&newEntity).Error
		return err
	} else if count > 0 {
		//update
	}

	return err
}
