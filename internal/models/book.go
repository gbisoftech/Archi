package models

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID        string    `form:"id"`
	Title     string    `form:"title"`
	Author    string    `form:"author"`
	Quantity  int       `form:"quantity"`
	UserID    string    `form:"user_id" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `form:"created_at"`
	UpdatedAt time.Time `form:"updated_at"`
}

func (b *Book) BeforeCreate(c *gin.Context) error {
	b.ID = uuid.New().String()
	userID, exists := c.Get("userID")
	if !exists {
		return errors.New("userID not found in context")
	}
	b.UserID = userID.(string)
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	return nil
}

func (b *Book) Create() error {

	result := DB.Create(&b)

	if result.Error != nil {
		log.Fatal(result.Error)
		return result.Error
	}

	return nil

}

func GetBooks() ([]Book, error) {
	var books []Book

	result := DB.Find(&books)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return books, nil

}

func GetBook(id string) (*Book, error) {
	var book Book

	result := DB.Where("id = ?", id).First(&book)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}
	return &book, nil
}

func (b *Book) Update(id string) error {

	result := DB.Model(&Book{}).Where("id = ?", id).Updates(b)

	if result.Error != nil {
		log.Fatal(result.Error)
		return result.Error
	}
	return nil
}

func DeleteBook(id string) error {

	result := DB.Delete(&Book{}, "id = ?", id)

	if result.Error != nil {
		log.Fatal(result.Error)
		return result.Error
	}
	return nil
}
