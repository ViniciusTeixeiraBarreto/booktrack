package author

import (
	service "booktrack/internal/service/author"
	"booktrack/models"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler struct {
	authorService service.AuthorService
}

func NewHandler() Handler {
	return Handler{
		authorService: service.NewAuthorService(),
	}
}

func (bh Handler) Create(c *fiber.Ctx) error {
	ctx := c.Context()

	var entity models.Author

	err := c.BodyParser(&entity)
	if err != nil {
		return c.Status(400).JSON(gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

	}

	entity, _ = bh.authorService.Create(ctx, entity)

	return c.Status(201).JSON(entity)
}

func (bh Handler) Get(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	newid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(gin.H{
			"error": "ID has to be integer",
		})

	}

	books, _ := bh.authorService.Get(ctx, newid)

	return c.Status(200).JSON(books)
}

func (bh Handler) GetAll(c *fiber.Ctx) error {
	ctx := c.Context()

	books, _ := bh.authorService.GetAll(ctx)

	return c.Status(200).JSON(books)
}

func (bh Handler) Update(c *fiber.Ctx) error {
	ctx := c.Context()
	var entity models.Author

	err := c.BodyParser(&entity)
	if err != nil {
		return c.Status(400).JSON(gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

	}

	books, _ := bh.authorService.Update(ctx, entity)

	return c.Status(200).JSON(books)
}

func (bh Handler) Delete(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	newid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(gin.H{
			"error": "ID has to be integer",
		})

	}

	_ = bh.authorService.Delete(ctx, newid)

	return c.Status(200).JSON(gin.H{
		"ok": true,
	})
}
