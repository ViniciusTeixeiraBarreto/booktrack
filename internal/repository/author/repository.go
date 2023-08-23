package author

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

func (r Repository) Create(ctx context.Context, entity models.Author) (models.Author, error) {
	db := database.GetConnection(ctx)

	err := db.Create(&entity).Error
	if err != nil {
		return entity, err
	}

	return entity, nil
}

func (r Repository) Get(ctx context.Context, id uuid.UUID) (models.Author, error) {
	db := database.GetConnection(ctx)

	var entity models.Author
	err := db.First(&entity, id).Error
	if err != nil {
		return entity, err
	}

	return entity, nil
}

func (r Repository) GetAll(ctx context.Context) ([]models.Author, error) {
	db := database.GetConnection(ctx)

	var entity []models.Author
	err := db.Find(&entity).Error
	if err != nil {
		return entity, err
	}

	return entity, nil
}

func (r Repository) Update(ctx context.Context, entity models.Author) (models.Author, error) {
	db := database.GetConnection(ctx)

	err := db.Save(&entity).Error
	if err != nil {
		return entity, err
	}

	return entity, nil
}

func (r Repository) Delete(ctx context.Context, id uuid.UUID) error {
	db := database.GetConnection(ctx)

	err := db.Delete(&models.Author{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
