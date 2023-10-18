package routes

import (
	handler "booktrack/internal/handler/author"

	"github.com/gofiber/fiber/v2"
)

func AuthorRouter(router fiber.Router) {
	authorHandler := handler.NewHandler()

	router.Get("/:id", authorHandler.Get)
	router.Get("/", authorHandler.GetAll)
	router.Post("/", authorHandler.Create)
	router.Put("/", authorHandler.Update)
	router.Delete("/:id", authorHandler.Delete)
}
