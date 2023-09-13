package http

import (
	controller "github.com/S4mkiel/p-1/infra/http/controllers"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"http",
	FiberModule,
	fx.Provide(controller.NewUserController),
	fx.Provide(controller.NewSexoController),
	fx.Invoke(RegisterControllers),
)

func RegisterControllers(app *fiber.App, userController *controller.UserController, sexoController *controller.SexoController) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	userController.RegisterRoutes(v1)
	sexoController.RegisterRoutes(v1)
}
