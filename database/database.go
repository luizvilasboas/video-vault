package database

import (
	"fmt"
	"log"
	"time"

	"gitlab.com/olooeez/video-vault/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func Connect() {
	dsn := "host=localhost user=root password=root dbname=videoVault port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("erro ao conectar ao banco de dados: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("erro ao obter o objeto DB: %v", err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	if err := runMigrations(db); err != nil {
		log.Fatalf("erro ao executar as migrações: %v", err)
	}

	baseConfig(db)

	DB = db
}

func ConnectForTest() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		log.Fatalf("erro ao conectar ao banco de dados: %v", err)
	}

	if err := runMigrations(db); err != nil {
		log.Fatalf("erro ao executar as migrações: %v", err)
	}

	baseConfig(db)

	DB = db
}

func CloseForTest() {
	if DB != nil {
		sqlDB, err := DB.DB()

		if err != nil {
			panic("failed to get test database connection")
		}

		sqlDB.Close()
	}
}

func runMigrations(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Video{}, &models.Category{}); err != nil {
		return fmt.Errorf("erro ao realizar migrações: %v", err)
	}
	return nil
}

func baseConfig(db *gorm.DB) {
	if err := db.FirstOrCreate(&models.Category{
		Title: "Livre",
		Color: "#FFF",
	}).Error; err != nil {
		log.Fatalf("Erro ao criar ou obter categoria padrão: %v", err)
	}
}
