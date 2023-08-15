package handler

import (
	"web-api/controller"
	"web-api/models"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return
	}

	book, _ = controller.Create(book)

	c.JSON(201, book)
}
