package handler

import (
	"booktrack/controller"
	"booktrack/models"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	ctx := c.Request.Context()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return
	}

	book, _ = controller.Create(ctx, book)

	c.JSON(201, book)
}
