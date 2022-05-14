package routes

import (
	"web-api/controller"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", controller.ShowBook)
			books.GET("/", controller.ShowBooks)
			books.POST("/", controller.CreateBook)
			books.PUT("/", controller.UpdateBooks)
			books.DELETE("/:id", controller.DeleteBook)
		}
	}

	return router
}
