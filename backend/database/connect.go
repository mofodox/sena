package database

import (
	"fmt"
	"log"
	"os"

	"github.com/mofodox/sena/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (err error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Singapore", os.Getenv("DBHOST"), os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBNAME"), os.Getenv("DBPORT"))

	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		return err
	}

	log.Println("Successfully connected to the database")

	DB.AutoMigrate(&models.Todo{})
	log.Println("DB migrated")

	return nil
}
