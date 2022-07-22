package database

import (
	"fmt"
	"log"
	"time"

	"github.com/vagnerpelais/napp-agenda/database/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func StartDB(connection_string string) {
	str := connection_string

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})

	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	}

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
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
