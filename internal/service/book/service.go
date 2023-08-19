package book

import (
	"booktrack/interfaces"
	"booktrack/internal/repository/book"
	"booktrack/models"
	"booktrack/models/services"
	"context"
	"errors"
	"fmt"
)

type BookService struct {
	bookRepository interfaces.BookRepository
}

func NewBookService(container *services.ServicesContainer) BookService {
	return BookService{
		bookRepository: book.NewRepository(),
	}
}

func (c BookService) Create(ctx context.Context, newBook models.Book) (models.Book, error) {
	newBook, err := c.bookRepository.Create(ctx, newBook)
	if err != nil {

		return newBook, errors.New(fmt.Sprintf("cannot create book: " + err.Error()))
	}

	return newBook, err
}
