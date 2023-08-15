package models

type Book struct {
	Metadata

	Name        string  `json:"name"`
	Description string  `json:"description"`
	MediumPrice float32 `json:"medium_price"`
	ImageURL    string  `json:"img_url"`

	Authors []Author `gorm:"many2many:book_authors;"`
}
