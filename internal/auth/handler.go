package auth

import (
	"github.com/gofiber/fiber/v2"
)

type UserDTO struct {
	Username   string
	Password   string
	IsApproved bool
}

type AuthHandler struct {
	authService AuthService
}

func NewHandler(authService *AuthService) *AuthHandler {
	return &AuthHandler{
		authService: *authService,
	}
}

func (h AuthHandler) Register(c *fiber.Ctx) error {
	return nil
}
