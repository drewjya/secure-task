package auth_controller

import (
	auth_service "github.com/drewjya/secure-task/service/auth"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	LoginController(c *fiber.Ctx) error
	RegisterController(c *fiber.Ctx) error
}

type AuthControllerImpl struct {
	service auth_service.AuthService
}

func NewAuthController(service auth_service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		service: service,
	}
}
