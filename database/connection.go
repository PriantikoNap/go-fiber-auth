package database

import (
	"github.com/PriantikoNap/go-fiber-auth.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=1 dbname=go-fiber-jwt port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cannot connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})

}
