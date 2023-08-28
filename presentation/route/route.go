package route

import (
	auth_controller "github.com/drewjya/secure-task/presentation/controller/auth"
	"github.com/drewjya/secure-task/repository"
	auth_service "github.com/drewjya/secure-task/service/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Setup(app *fiber.App, queries *repository.Queries) {
	authService := auth_service.NewAuthService(queries)

	authController := auth_controller.NewAuthController(authService)
	app.Use(recover.New())
	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"data": "response",
		})
	})
	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "Not Implemented",
		})
	})
	v1 := api.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "Not Implemented",
		})
	})
	auth := v1.Group("/auth")
	auth.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "Not Implemented",
		})
	})
	auth.Post("/login", authController.LoginController)
	auth.Post("/signup", authController.RegisterController)

}
