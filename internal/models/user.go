package models

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `form:"id"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (u *User) BeforeCreate() error {
	u.ID = uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) Create() error {

	result := DB.Create(&u)

	if result.Error != nil {
		log.Fatal(result.Error)
		return result.Error
	}

	return nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User

	result := DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
