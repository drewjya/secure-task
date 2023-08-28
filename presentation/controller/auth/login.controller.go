package auth_controller

import (
	domain "github.com/drewjya/secure-task/domain/auth"
	"github.com/drewjya/secure-task/domain/response"
	error_handler "github.com/drewjya/secure-task/libs/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (controller *AuthControllerImpl) LoginController(c *fiber.Ctx) error {
	loginDto := new(domain.LoginDTO)

	if err := c.BodyParser(loginDto); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(response.NewError(fiber.StatusBadRequest, fiber.Map{
			"data": err.Error(),
		}))
	}

	validate := validator.New()
	if err := validate.Struct(loginDto); err != nil {
		errors := error_handler.CustomErrorHandler(err)
		return c.Status(fiber.StatusBadRequest).JSON(response.NewError(fiber.StatusBadRequest, errors))
	}

	session, err := controller.service.LoginService(loginDto)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(response.NewError(fiber.StatusNotFound, fiber.Map{
			"data": err.Error(),
		}))
	}

	return c.JSON(response.NewResponse(fiber.StatusAccepted, fiber.Map{
		"session": session,
	}))

}
