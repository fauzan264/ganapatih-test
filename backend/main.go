package main

import (
	"fmt"
	"log"

	"github.com/fauzan264/backend/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// db := config.InitDatabase()
	router := fiber.New()
	router.Use(cors.New())

	if err := router.Listen(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)); err != nil {
		log.Println("Error: ", err)
	}
}