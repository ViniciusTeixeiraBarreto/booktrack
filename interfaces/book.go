package interfaces

import (
	"context"
	"web-api/models"
)

type BookController interface {
	Create(ctx context.Context, book models.Book) (models.Book, error)
}

type BookRepository interface {
	Create(ctx context.Context, book models.Book) (models.Book, error)
}

type BookHandler interface{}
