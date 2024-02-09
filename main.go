package main

import (
	"gb-backend/handler"
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
	var PORT string = os.Getenv("PORT")

	app := fiber.New()

	authHandler := handler.AuthHandler{}

	app.Post("/auth/register", authHandler.HandleRegister)

	app.Listen(PORT)
}
