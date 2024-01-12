package database

import (
	"gitlab.com/olooeez/video-vault/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	connection := "host=localhost user=root password=root dbname=videoVault port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(connection))
	if err != nil {
		panic("could not connect to the database")
	}

	DB.AutoMigrate(&models.Video{})
}
