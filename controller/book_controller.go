package controller

import (
	"errors"
	"fmt"
	"strconv"
	"web-api/database"
	"web-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(newBook models.Book) (models.Book, error) {
	db := database.GetDatabase()

	err := db.Create(&newBook).Error
	if err != nil {

		return newBook, errors.New(fmt.Sprintf("cannot create book: " + err.Error()))
	}

	return newBook, err
}

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book

	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list books: " + err.Error(),
		})

		return
	}

	c.JSON(200, books)
}

func UpdateBooks(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return
	}

	err = db.Save(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.Delete(&book, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find book: " + err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"ok": true,
	})
}

func CountBooks(c *gin.Context) {
	db := database.GetDatabase()

	var countBooks int64

	err := db.Find(&models.Book{}).Count(&countBooks).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list books: " + err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"books": countBooks,
	})
}

func SearchesBooks(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book
	var books []models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return
	}

	err = db.Find(&books, book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find books: " + err.Error(),
		})

		return
	}

	c.JSON(200, books)
}

func ChangeMediumPriceBook(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var objPrice struct {
		MediumPrice float32 `json:"medium_price"`
	}

	err = c.ShouldBindJSON(&objPrice)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return
	}

	var book models.Book
	err = db.First(&book, newid).UpdateColumns(models.Book{
		MediumPrice: objPrice.MediumPrice,
	}).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update medium price of this book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

//TODO remove from this file
func MediumPriceGreaterThan(db *gorm.DB) *gorm.DB {
	return db.Where("medium_price > ?", 18)
}

func MediumPriceLessThan(db *gorm.DB) *gorm.DB {
	return db.Where("medium_price < ?", 20)
}

func FilterBetweenMediumPriceBook(c *gin.Context) {
	db := database.GetDatabase()

	var objPrice struct {
		FirstValue  float32 `json:"first_value"`
		SecondValue float32 `json:"second_value"`
	}

	err := c.ShouldBindJSON(&objPrice)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find JSON: " + err.Error(),
		})

		return
	}

	var books []models.Book
	err = db.Where("medium_price > ? and medium_price < ?", objPrice.FirstValue, objPrice.SecondValue).Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update medium price of this book: " + err.Error(),
		})

		return
	}

	c.JSON(200, books)
}
