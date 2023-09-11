package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ImanAski/janotan-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tehran",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
