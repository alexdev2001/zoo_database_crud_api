package database_connect

import (
	"alexdev2001/zoo_database_crud_api/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func DBConnection() {
	// load the env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// obtain attributes from the .env
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")

	fmt.Println(db_host, db_name, db_port, db_user, db_password)

	// format a dsn string to be used to connect to database
	dsn := fmt.Sprintf("dbname=%s host=%s user=%s password=%s  port=%s", db_name, db_host, db_user, db_password, db_port)

	// connect to the database using the dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	// auto migrate the database model
	err = db.AutoMigrate(&models.Animal{})
	if err != nil {
		log.Fatal("Error creating animal table")
	}

}
