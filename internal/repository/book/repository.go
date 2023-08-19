package book

import (
	"booktrack/models"
	"booktrack/pkg/database"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Repository struct {
}

func NewRepository() Repository {
	return Repository{}
}

func (r Repository) Create(ctx context.Context, book models.Book) (models.Book, error) {
	db := database.GetConnection(ctx)

	err := db.Create(&book).Error
	if err != nil {
		return book, errors.New(fmt.Sprintf("cannot create book: " + err.Error()))
	}

	return book, nil
}

func (r Repository) Get(ctx context.Context, bookID uuid.UUID) (models.Book, error) {
	db := database.GetConnection(ctx)

	var book models.Book
	err := db.First(&book, bookID).Error
	if err != nil {
		return book, errors.New(fmt.Sprintf("cannot find book: " + err.Error()))
	}

	return book, nil
}

func (r Repository) GetAll(ctx context.Context) ([]models.Book, error) {
	db := database.GetConnection(ctx)

	var book []models.Book
	err := db.Find(&book).Error
	if err != nil {
		return book, errors.New(fmt.Sprintf("cannot find all books: " + err.Error()))
	}

	return book, nil
}

func (r Repository) Update(ctx context.Context, bookToUpdate models.Book) (models.Book, error) {
	db := database.GetConnection(ctx)

	err := db.Save(&bookToUpdate).Error
	if err != nil {
		return bookToUpdate, errors.New(fmt.Sprintf("cannot find all books: " + err.Error()))
	}

	return bookToUpdate, nil
}

func (r Repository) Delete(ctx context.Context, id uuid.UUID) error {
	db := database.GetConnection(ctx)

	err := db.Delete(&models.Book{}, id).Error
	if err != nil {
		return errors.New(fmt.Sprintf("cannot find all books: " + err.Error()))
	}

	return nil
}

func (r Repository) Count(ctx context.Context) (int64, error) {
	db := database.GetConnection(ctx)

	var countBooks int64

	err := db.Find(&models.Book{}).Count(&countBooks).Error
	if err != nil {
		return countBooks, err
	}

	return countBooks, nil
}

func (r Repository) Search(ctx context.Context, search models.Book) ([]models.Book, error) {
	db := database.GetConnection(ctx)

	books := []models.Book{}

	err := db.Find(&books, search).Error
	if err != nil {
		return books, err
	}

	return books, nil
}

func (r Repository) FilterBetweenMediumPriceBook(ctx context.Context, firstValue float64, secondValue float64) ([]models.Book, error) {
	db := database.GetConnection(ctx)

	var books []models.Book
	err := db.Where("medium_price > ? and medium_price < ?", firstValue, secondValue).Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
}
