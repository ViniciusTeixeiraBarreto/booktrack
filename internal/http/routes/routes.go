package routes

import (
	handler "booktrack/internal/handler/book"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		bookHandler := handler.NewHandler()

		books := main.Group("books")
		{
			books.GET("/:id", bookHandler.Get)
			books.GET("/", bookHandler.GetAll)
			books.POST("/", bookHandler.Create)
			books.PUT("/", bookHandler.Update)
			books.DELETE("/:id", bookHandler.Delete)

			books.GET("/count", bookHandler.CountBooks)
			books.POST("/searches", bookHandler.SearchesBooks)
			books.PUT("/:id/average-price", bookHandler.ChangeAveragePrice)
			books.POST("/sale", bookHandler.FilterBetweenMediumPriceBook)

		}
	}

	return router
}
