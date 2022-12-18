package database

import (
	"fmt"
	"github.com/valianx/gin-template/pkg/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

var DB *gorm.DB

func ConnectDataBase() {
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3000"
	}
	portInt, err := strconv.Atoi(port)

	connect := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		os.Getenv("DB_HOST"), portInt, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  connect,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		return
	}
}
