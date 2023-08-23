package book

import (
	"booktrack/models"
	"booktrack/pkg/database"
	"context"

	"github.com/google/uuid"
)

type Repository struct {
}

func NewRepository() Repository {
	return Repository{}
}

func (r Repository) Create(ctx context.Context, entity models.Book) (models.Book, error) {
	db := database.GetConnection(ctx)

	err := db.Create(&entity).Error
	if err != nil {
		return entity, err
	}

	return entity, nil
}

func (r Repository) Get(ctx context.Context, id uuid.UUID) (models.Book, error) {
	db := database.GetConnection(ctx)

	var book models.Book
	err := db.First(&book, id).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r Repository) GetAll(ctx context.Context) ([]models.Book, error) {
	db := database.GetConnection(ctx)

	var book []models.Book
	err := db.Find(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r Repository) Update(ctx context.Context, entity models.Book) (models.Book, error) {
	db := database.GetConnection(ctx)

	err := db.Save(&entity).Error

	if err != nil {
		return entity, err
	}

	return entity, nil
}

func (r Repository) Delete(ctx context.Context, id uuid.UUID) error {
	db := database.GetConnection(ctx)
	err := db.Delete(&models.Book{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) Count(ctx context.Context) (int64, error) {
	db := database.GetConnection(ctx)
	var count int64

	err := db.Find(&models.Book{}).Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}

func (r Repository) Search(ctx context.Context, search models.Book) ([]models.Book, error) {
	db := database.GetConnection(ctx)
	entities := []models.Book{}

	err := db.Find(&entities, search).Error
	if err != nil {
		return entities, err
	}

	return entities, nil
}

func (r Repository) FilterBetweenMediumPriceBook(ctx context.Context, firstValue float64, secondValue float64) ([]models.Book, error) {
	db := database.GetConnection(ctx)
	var entities []models.Book

	err := db.Where("medium_price > ? and medium_price < ?", firstValue, secondValue).Find(&entities).Error
	if err != nil {
		return entities, err
	}

	return entities, nil
}
