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
			books.GET("/:id", bookHandler.ShowBook)
			books.GET("/", bookHandler.ShowBooks)
			books.POST("/", bookHandler.Create)
			books.PUT("/", bookHandler.UpdateBooks)
			books.DELETE("/:id", bookHandler.DeleteBook)

			books.GET("/count", bookHandler.CountBooks)
			books.POST("/searches", bookHandler.SearchesBooks)
			books.PUT("/:id/mediumPrice", bookHandler.ChangeMediumPriceBook)
			books.POST("/sale", bookHandler.FilterBetweenMediumPriceBook)

		}
	}

	return router
}
