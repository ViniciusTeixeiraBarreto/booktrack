package book

import (
	"booktrack/internal/repository/book"
	"booktrack/models"
	"context"

	"github.com/google/uuid"
)

type BookService struct {
	bookRepository book.Repository
}

func NewBookService() BookService {
	return BookService{
		bookRepository: book.NewRepository(),
	}
}

func (c BookService) Create(ctx context.Context, newBook models.Book) (models.Book, error) {
	return c.bookRepository.Create(ctx, newBook)
}

func (c BookService) Get(ctx context.Context, id uuid.UUID) (models.Book, error) {
	return c.bookRepository.Get(ctx, id)
}

func (c BookService) GetAll(ctx context.Context) ([]models.Book, error) {
	return c.bookRepository.GetAll(ctx)
}

func (c BookService) Update(ctx context.Context, bookToUpdate models.Book) (models.Book, error) {
	return c.bookRepository.Update(ctx, bookToUpdate)
}

func (c BookService) Delete(ctx context.Context, id uuid.UUID) error {
	return c.bookRepository.Delete(ctx, id)
}

func (c BookService) Count(ctx context.Context) (int64, error) {
	return c.bookRepository.Count(ctx)
}

func (c BookService) Search(ctx context.Context, search models.Book) ([]models.Book, error) {
	return c.bookRepository.Search(ctx, search)
}

func (c BookService) ChangeMediumPriceBook(ctx context.Context, mediumPrice float32, bookId uuid.UUID) (models.Book, error) {
	return c.bookRepository.Update(ctx, models.Book{
		Metadata:    models.Metadata{ID: bookId},
		MediumPrice: mediumPrice,
	})
}

func (c BookService) FilterBetweenMediumPriceBook(ctx context.Context, firstValue float64, secondValue float64) ([]models.Book, error) {
	return c.bookRepository.FilterBetweenMediumPriceBook(ctx, firstValue, secondValue)
}
