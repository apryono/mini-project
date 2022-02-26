package bootstrap

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mini-project/server/bootstrap/routers"
	"github.com/mini-project/server/handlers"
)

//RegisterRouters use to register all routes
func (boot Bootstrap) RegisterRouters() {
	handler := handlers.Handler{
		FiberApp:   boot.App,
		ContractUC: &boot.ContractUC,
		Validator:  boot.Validator,
		Translator: boot.Translator,
	}

	// Testing
	boot.App.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON("success")
	})

	apiV1 := boot.App.Group("/v1")

	cakeRoute := routers.CakesRoutes{RouterGroup: apiV1, Handler: handler}
	cakeRoute.RegisterRoute()

}
