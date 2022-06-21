package main

import (
	"github.com/joho/godotenv"
	"github.com/luccasalves/newsletter-go/server"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	s := server.CreateServer()
	s.Init()
}
