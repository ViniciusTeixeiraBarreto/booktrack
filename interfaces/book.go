package interfaces

import (
	"booktrack/models"
	"context"
)

type BookController interface {
	Create(ctx context.Context, book models.Book) (models.Book, error)
}

type BookRepository interface {
	Create(ctx context.Context, book models.Book) (models.Book, error)
}

type BookHandler interface{}
