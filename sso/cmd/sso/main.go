package main

import (
	"fmt"
	"log"
	"sso/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../../local.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: инициализировать логгер

	// TODO: инициализировать приложение (app)
}
