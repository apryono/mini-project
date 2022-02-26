package routers

import (
	"github.com/mini-project/server/handlers"

	"github.com/gofiber/fiber/v2"
)

type CakesRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

func (c CakesRoutes) RegisterRoute() {

}
