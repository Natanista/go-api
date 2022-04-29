package database

import (
	"fmt"
	"log"
	// "os"
	"time"

	"github.com/natanista/go-api/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	// host := os.Getenv("DATABASE_HOST")
	// dialect := os.Getenv("DATABASE_DIALECT")
	// port := os.Getenv("DATABASE_PORT")
	// user := os.Getenv("DATABASE_USER")
	// dbname := os.Getenv("DATABASE_DBNAME")
	// password := os.Getenv("DATABASE_PASSWORD")

	// str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)

	
	database, err := gorm.Open(postgres.Open("postgres://vtvsbfqobfssqw:edee9aa648b04a7940144994afe1925a1c83fdcf876689394759e272f4207fe5@ec2-34-194-73-236.compute-1.amazonaws.com:5432/d2n1bq3i9rb6vt"), &gorm.Config{})

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
