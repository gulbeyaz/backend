package main

import (
	"context"
	"gb-backend/internal/auth"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	PORT := os.Getenv("PORT")
	MONGO_URL := os.Getenv("MONGO_URL")

	APP := fiber.New()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MONGO_URL))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	authRespository := auth.NewRepository("")
	authService := auth.NewService(authRespository)
	authHandler := auth.NewHandler(authService)

	APP.Post("/auth/register", authHandler.Register)

	log.Printf("Server started on port %s", PORT)
	log.Fatal(APP.Listen(":" + PORT))
}
