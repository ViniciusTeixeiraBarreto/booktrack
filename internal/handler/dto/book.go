package dto

import "booktrack/models"

type CreateBook struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	MediumPrice float32         `json:"medium_price"`
	ImageURL    string          `json:"img_url"`
	Authors     []models.Author `json:"authors"`
}

func (cb CreateBook) ToBook() models.Book {
	return models.Book{
		Name:        cb.Name,
		Description: cb.Description,
		MediumPrice: cb.MediumPrice,
		ImageURL:    cb.ImageURL,
		Authors:     cb.Authors,
	}
}
