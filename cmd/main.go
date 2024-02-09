package main

import (
	"gb-backend/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	authHandler := handler.AuthHandler{}

	app.POST("/auth/register", authHandler.HandleRegister)

	app.Start(":8081")
}
