package repository

import (
	"context"
	"errors"
	"fmt"
	"web-api/models"
	"web-api/pkg/database"
)

type Repository struct {
}

func NewRepository() {

}

func (r Repository) Create(ctx context.Context, book models.Book) (models.Book, error) {
	db := database.GetDatabase()

	err := db.Create(&book).Error
	if err != nil {
		return book, errors.New(fmt.Sprintf("cannot create book: " + err.Error()))
	}

	return book, nil
}
