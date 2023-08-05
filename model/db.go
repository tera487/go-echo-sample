package model

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB は
var DB *gorm.DB
var err error

func SetupDB() {
	dsn := "tester:password@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(dsn + "database can't connect")
	}
	err := DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatalln(dsn + "database can't connect")
	}
}
