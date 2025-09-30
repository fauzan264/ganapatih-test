package main

import (
	"fmt"
	"log"

	"github.com/fauzan264/backend/config"
	"github.com/fauzan264/backend/dto/response"
	"github.com/fauzan264/backend/handlers"
	"github.com/fauzan264/backend/middleware"
	repository "github.com/fauzan264/backend/repositories"
	service "github.com/fauzan264/backend/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	db := config.InitDatabase()
	router := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			msg := err.Error()

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				msg = e.Message
			}

			return c.Status(code).JSON(response.Response{
				Message: msg,
			})
		},
	})
	
	router.Use(cors.New())
	router.Use(logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
	}))

	// repositories
	authRepository := repository.NewAuthRepository(db)

	// services
	authService := service.NewAuthService(authRepository)

	// handlers
	authHandler :=  handlers.NewAuthHandler(authService)

	authMiddleware := middleware.AuthMiddleware(authService)
	api := router.Group("/api")

	// auth
	api.Post("/register", authHandler.RegisterUser)
	api.Post("/login", authHandler.LoginUser)
	api.Get("/session", authMiddleware, authHandler.SessionUser)


	if err := router.Listen(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)); err != nil {
		log.Println("Error: ", err)
	}
}