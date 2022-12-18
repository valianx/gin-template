package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func GetPort() (n int, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	n, err = strconv.Atoi(port)
	if err != nil {
		fmt.Printf("Error converting string to int: %v\n", err)
		return 0, err
	}
	return n, nil
}
