package handler

import (
	"booktrack/internal/service/book"
	"booktrack/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	bookService book.BookService
}

func NewHandler() Handler {
	return Handler{
		bookService: book.NewBookService(),
	}
}

func (bh Handler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var newBook models.Book

	err := c.ShouldBindJSON(&newBook)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return
	}

	newBook, _ = bh.bookService.Create(ctx, newBook)

	c.JSON(201, newBook)
}

func (bh Handler) ShowBook(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	newid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	books, _ := bh.bookService.Get(ctx, newid)

	c.JSON(200, books)
}

func (bh Handler) ShowBooks(c *gin.Context) {
	ctx := c.Request.Context()

	books, _ := bh.bookService.GetAll(ctx)

	c.JSON(200, books)
}

func (bh Handler) UpdateBooks(c *gin.Context) {
	ctx := c.Request.Context()
	var bookToUpdate models.Book

	err := c.ShouldBindJSON(&bookToUpdate)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return
	}

	books, _ := bh.bookService.Update(ctx, bookToUpdate)

	c.JSON(200, books)
}

func (bh Handler) DeleteBook(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	newid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	_ = bh.bookService.Delete(ctx, newid)

	c.JSON(200, gin.H{
		"ok": true,
	})
}

func (bh Handler) CountBooks(c *gin.Context) {
	ctx := c.Request.Context()
	var countBooks int64

	countBooks, _ = bh.bookService.Count(ctx)

	c.JSON(200, gin.H{
		"books": countBooks,
	})
}

func (bh Handler) SearchesBooks(c *gin.Context) {
	ctx := c.Request.Context()
	var book models.Book
	var books []models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return
	}

	books, _ = bh.bookService.Search(ctx, book)

	c.JSON(200, books)
}

func (bh Handler) ChangeMediumPriceBook(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	newid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	var objPrice struct {
		MediumPrice float32 `json:"medium_price"`
	}

	err = c.ShouldBindJSON(&objPrice)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return
	}

	books, _ := bh.bookService.ChangeMediumPriceBook(ctx, objPrice.MediumPrice, newid)

	c.JSON(200, books)
}

func (bh Handler) FilterBetweenMediumPriceBook(c *gin.Context) {
	ctx := c.Request.Context()

	var objPrice struct {
		FirstValue  float64 `json:"first_value"`
		SecondValue float64 `json:"second_value"`
	}

	err := c.ShouldBindJSON(&objPrice)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return
	}

	books, _ := bh.bookService.FilterBetweenMediumPriceBook(ctx, objPrice.FirstValue, objPrice.SecondValue)

	c.JSON(200, books)
}
