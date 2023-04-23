package main

import (
	"SOC_N5_14_BTL/cmd/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//cli.Migrate()
	server.Start()
}
