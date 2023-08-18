package handler

import (
	"booktrack/controller/book"
	"booktrack/models"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	ctx := c.Request.Context()

	var newBook models.Book

	err := c.ShouldBindJSON(&newBook)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return
	}

	controller := book.NewController()

	newBook, _ = controller.Create(ctx, newBook)

	c.JSON(201, newBook)
}
