package main

import (
	"fmt"
	"github.com/Mogza/FandC/pkg/db"
	"github.com/Mogza/FandC/pkg/handlers"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file : %v", err)
	}
	fmt.Println("First Test")
	DB := db.Init()
	h := handlers.New(DB)
	if h.DB != nil {
		fmt.Println("DB Successfully Initialized")
	}
	return
}
