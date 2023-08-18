package routes

import (
	"booktrack/internal/handler"
	service "booktrack/internal/service"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", service.ShowBook)
			books.GET("/", service.ShowBooks)
			books.POST("/", handler.Create)
			books.PUT("/", service.UpdateBooks)
			books.DELETE("/:id", service.DeleteBook)

			books.GET("/count", service.CountBooks)
			books.POST("/searches", service.SearchesBooks)
			books.PUT("/:id/mediumPrice", service.ChangeMediumPriceBook)
			books.POST("/sale", service.FilterBetweenMediumPriceBook)

		}
	}

	return router
}
