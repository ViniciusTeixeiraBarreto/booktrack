package handler

import (
	book_service "booktrack/internal/service/book"
	"booktrack/models"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler struct {
	bookService book_service.BookService
}

func NewHandler() Handler {
	return Handler{
		bookService: book_service.NewBookService(),
	}
}

func (bh Handler) Create(c *fiber.Ctx) error {
	ctx := c.Context()

	var newBook models.Book

	err := c.BodyParser(&newBook)
	if err != nil {
		c.Status(400).JSON(gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return nil
	}

	newBook, _ = bh.bookService.Create(ctx, newBook)

	return c.Status(201).JSON(newBook)
}

func (bh Handler) Get(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	newid, err := uuid.Parse(id)
	if err != nil {
		c.Status(400).JSON(gin.H{
			"error": "ID has to be integer",
		})

		return nil
	}

	books, _ := bh.bookService.Get(ctx, newid)

	return c.Status(200).JSON(books)
}

func (bh Handler) GetAll(c *fiber.Ctx) error {
	ctx := c.Context()

	books, _ := bh.bookService.GetAll(ctx)

	return c.Status(200).JSON(books)
}

func (bh Handler) Update(c *fiber.Ctx) error {
	ctx := c.Context()
	var bookToUpdate models.Book

	err := c.BodyParser(&bookToUpdate)
	if err != nil {
		c.Status(400).JSON(gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return nil
	}

	books, _ := bh.bookService.Update(ctx, bookToUpdate)

	return c.Status(200).JSON(books)
}

func (bh Handler) Delete(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	newid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(gin.H{
			"error": "ID has to be uuid",
		})
	}

	if err = bh.bookService.Delete(ctx, newid); err != nil {
		return c.Status(400).JSON(gin.H{
			"error": "broken service",
		})
	}

	return c.Status(200).JSON(gin.H{
		"ok": true,
	})
}

func (bh Handler) CountBooks(c *fiber.Ctx) error {
	ctx := c.Context()
	var countBooks int64

	countBooks, _ = bh.bookService.Count(ctx)

	return c.Status(200).JSON(gin.H{
		"books": countBooks,
	})
}

func (bh Handler) SearchesBooks(c *fiber.Ctx) error {
	ctx := c.Context()
	var book models.Book
	var books []models.Book

	err := c.BodyParser(&book)
	if err != nil {
		c.Status(400).JSON(gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return nil
	}

	books, _ = bh.bookService.Search(ctx, book)

	return c.Status(200).JSON(books)
}

func (bh Handler) ChangeAveragePrice(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	newid, err := uuid.Parse(id)
	if err != nil {
		c.Status(400).JSON(gin.H{
			"error": "ID has to be integer",
		})

		return nil
	}

	var objPrice struct {
		MediumPrice float32 `json:"medium_price"`
	}

	err = c.BodyParser(&objPrice)
	if err != nil {
		c.Status(400).JSON(gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return nil
	}

	books, _ := bh.bookService.ChangeAveragePrice(ctx, objPrice.MediumPrice, newid)

	return c.Status(200).JSON(books)
}

func (bh Handler) FilterBetweenMediumPriceBook(c *fiber.Ctx) error {
	ctx := c.Context()

	var objPrice struct {
		FirstValue  float64 `json:"first_value"`
		SecondValue float64 `json:"second_value"`
	}

	err := c.BodyParser(&objPrice)
	if err != nil {
		c.Status(400).JSON(gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return nil
	}

	books, _ := bh.bookService.FilterBetweenMediumPriceBook(ctx, objPrice.FirstValue, objPrice.SecondValue)

	return c.Status(200).JSON(books)
}
