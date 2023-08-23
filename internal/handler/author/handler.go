package author

import (
	service "booktrack/internal/service/author"
	"booktrack/models"

	"github.com/gin-gonic/gin"
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

func (bh Handler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var entity models.Author

	err := c.ShouldBindJSON(&entity)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return
	}

	entity, _ = bh.authorService.Create(ctx, entity)

	c.JSON(201, entity)
}

func (bh Handler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	newid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	books, _ := bh.authorService.Get(ctx, newid)

	c.JSON(200, books)
}

func (bh Handler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	books, _ := bh.authorService.GetAll(ctx)

	c.JSON(200, books)
}

func (bh Handler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var entity models.Author

	err := c.ShouldBindJSON(&entity)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return
	}

	books, _ := bh.authorService.Update(ctx, entity)

	c.JSON(200, books)
}

func (bh Handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	newid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	_ = bh.authorService.Delete(ctx, newid)

	c.JSON(200, gin.H{
		"ok": true,
	})
}
