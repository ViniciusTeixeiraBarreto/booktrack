package book

import "gorm.io/gorm"

func mediumPriceGreaterThan(db *gorm.DB) *gorm.DB {
	return db.Where("medium_price > ?", 18)
}

func mediumPriceLessThan(db *gorm.DB) *gorm.DB {
	return db.Where("medium_price < ?", 20)
}
