package routers

import (
	"github.com/mini-project/server/handlers"

	"github.com/gofiber/fiber/v2"
)

type CakesRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

func (route CakesRoutes) RegisterRoute() {
	handler := handlers.CakeHandler{Handler: route.Handler}

	r := route.RouterGroup.Group("/api/cake")
	r.Post("", handler.Add)
	r.Get("/id/:id", handler.DetailCake)
	r.Get("", handler.ListAllCake)
	r.Put("/id/:id", handler.Edit)
	r.Delete("/id/:id", handler.Delete)
}
