package book

import (
	"booktrack/interfaces"
	"booktrack/models"
	"booktrack/repository/book"
	"context"
	"errors"
	"fmt"
)

type Controller struct {
	bookRepository interfaces.BookRepository
}

func NewController() Controller {
	return Controller{
		bookRepository: book.NewRepository(),
	}
}

func (c Controller) Create(ctx context.Context, newBook models.Book) (models.Book, error) {
	newBook, err := c.bookRepository.Create(ctx, newBook)
	if err != nil {

		return newBook, errors.New(fmt.Sprintf("cannot create book: " + err.Error()))
	}

	return newBook, err
}
