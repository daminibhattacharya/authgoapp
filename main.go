package main

import (
	"fmt"
	"log"
	"auth-go-app/db"
	"auth-go-app/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	err := db.Init()
	if err != nil {
		fmt.Println("Failed to connect to database " + err.Error())
		return
	}
	defer db.Close()
	r := router.NewRouter()
	r.Run(":8080")
}
