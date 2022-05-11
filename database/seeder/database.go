package database

import (
	"fmt"
	"log"
	"os"

	// "os"

	// "os"
	"time"

	"github.com/natanista/go-api/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	// uri := os.Getenv("DATABASE_URL")

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_DBNAME")
	password := os.Getenv("DB_PASSWORD")
	
	//create url 
	str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)

	//docker postgres db
		// str := "host=localhost port=25431 user=admin dbname=books sslmode=disable password=123456"


	
	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	}

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigration(db)
}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
