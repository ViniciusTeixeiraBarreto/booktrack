package author

import (
	author_repository "booktrack/internal/repository/author"
	"booktrack/models"
	"context"

	"github.com/google/uuid"
)

type AuthorService struct {
	authorRepository author_repository.Repository
}

func NewAuthorService() AuthorService {
	return AuthorService{
		authorRepository: author_repository.NewRepository(),
	}
}

func (c AuthorService) Create(ctx context.Context, entity models.Author) (models.Author, error) {
	return c.authorRepository.Create(ctx, entity)
}

func (c AuthorService) Get(ctx context.Context, id uuid.UUID) (models.Author, error) {
	return c.authorRepository.Get(ctx, id)
}

func (c AuthorService) GetAll(ctx context.Context) ([]models.Author, error) {
	return c.authorRepository.GetAll(ctx)
}

func (c AuthorService) Update(ctx context.Context, entity models.Author) (models.Author, error) {
	return c.authorRepository.Update(ctx, entity)
}

func (c AuthorService) Delete(ctx context.Context, id uuid.UUID) error {
	return c.authorRepository.Delete(ctx, id)
}
