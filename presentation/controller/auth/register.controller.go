package auth_controller

import (
	"log"

	domain "github.com/drewjya/secure-task/domain/auth"
	"github.com/drewjya/secure-task/domain/response"
	"github.com/drewjya/secure-task/libs/hash"
	error_handler "github.com/drewjya/secure-task/libs/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (controller *AuthControllerImpl) RegisterController(c *fiber.Ctx) error {
	registerDto := new(domain.RegisterDTO)

	if err := c.BodyParser(registerDto); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(response.NewError(fiber.StatusBadRequest, err.Error()))
	}
	validate := validator.New()
	if err := validate.Struct(registerDto); err != nil {
		errors := error_handler.CustomErrorHandler(err)
		return c.Status(fiber.StatusBadRequest).JSON(response.NewError(fiber.StatusBadRequest, errors))
	}
	newPassword, err := hash.HashPassword(registerDto.Password)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(response.NewError(fiber.StatusBadRequest, err.Error()))

	}
	registerDto.Password = newPassword
	log.Println(registerDto.Password)
	session, err := controller.service.RegisterService(registerDto)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(response.NewError(fiber.StatusBadRequest, error_handler.SqlErrorHandler(err.Error())))
	}
	c.Status(fiber.StatusCreated)
	return c.JSON(response.NewResponse(fiber.StatusCreated, fiber.Map{
		"session": session,
	}))

}
