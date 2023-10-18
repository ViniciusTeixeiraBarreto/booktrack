package routes

import (
	handler "booktrack/internal/handler/book"

	"github.com/gofiber/fiber/v2"
)

func BookRouter(router fiber.Router) {
	bookHandler := handler.NewHandler()

	router.Get("/:id", bookHandler.Get)
	router.Get("/", bookHandler.GetAll)
	router.Post("/", bookHandler.Create)
	router.Put("/", bookHandler.Update)
	router.Delete("/:id", bookHandler.Delete)

	router.Get("/count", bookHandler.CountBooks)
	router.Post("/searches", bookHandler.SearchesBooks)
	router.Put("/:id/average-price", bookHandler.ChangeAveragePrice)
	router.Post("/sale", bookHandler.FilterBetweenMediumPriceBook)
}
