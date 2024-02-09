package main

import (
	"gb-backend/internal/auth"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	PORT := os.Getenv("PORT")
	DB := "mongo"

	APP := fiber.New()

	authRespository := auth.NewRepository(DB)
	authService := auth.NewService(authRespository)
	authHandler := auth.NewHandler(authService)

	APP.Post("/auth/register", authHandler.Register)

	APP.Listen(PORT)
}
