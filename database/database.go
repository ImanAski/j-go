package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ImanAski/janotan-api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("couldn't load the env file")
	}
	dsn := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tehran",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	log.Println("Connecting to the database...")
	if err != nil {
		log.Fatal("Failed Connecting to database", err)
	}

	log.Println("Connected to the Database")
	log.Println("Migrating Models...")
	db.AutoMigrate(&models.User{})

	DB = DbInstance{
		Db: db,
	}
}
