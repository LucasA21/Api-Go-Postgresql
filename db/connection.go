package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=lucas password=mysecretpassword dbname=gorm port=5434"

var DB *gorm.DB

func DBConeection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	} else {
		log.Println("DB connected")
	}
}
