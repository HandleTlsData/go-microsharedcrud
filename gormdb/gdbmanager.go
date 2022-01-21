package gdbmanager

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"sharedcrud/dbmanager"
	"sharedcrud/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var SqlDB *sql.DB

const StrNoRecords = "record not found"
const EntityAlreadyExists = "entity already exists"

func InitDB() {
	var err error
	var dsn string
	switch dbmanager.CurrentAppConfig {
	case "alpha":
		dsn = fmt.Sprintf("alphadb:kFimTXzzdnpTDFLE@tcp(149.34.51.91:3306)/alphadb?charset=utf8mb4&parseTime=True")
	case "beta":
		dsn = fmt.Sprintf("betadb:2KbB7r8J28B8CFyt@tcp(149.34.51.91:3306)/betadb?charset=utf8mb4&parseTime=True")
	default:
		log.Fatal("Unknown App Config. API Functions unimplemented")
	}

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

	if err != nil {
		if err.Error() == StrNoRecords {
			err = db.Model(&models.Entity{}).Create(&newEntity).Error
			return err
		}
	}

	return fmt.Errorf(EntityAlreadyExists)
}

func UpdateEntity(newEntity models.Entity, entityID int64) error {
	var user models.Entity
	var count int64
	err := db.Model(&models.Entity{}).Where(models.Entity{ID: entityID}).First(&user).Count(&count).Error

	if err != nil {
		return err
	}

	user.Description = newEntity.Description
	user.Name = newEntity.Name
	user.Status = newEntity.Status

	err = db.Save(&user).Error
	return err
}
