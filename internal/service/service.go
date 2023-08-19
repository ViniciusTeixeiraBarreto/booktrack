package service

import (
	"booktrack/internal/service/book"
	"booktrack/models/services"
)

func NewServiceContainer() services.ServicesContainer {
	container := services.ServicesContainer{}

	container.Book = book.NewBookService(&container)

	return container
}
