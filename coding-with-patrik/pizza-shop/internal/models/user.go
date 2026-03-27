package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id uint `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

type UserModel struct {
	DB *gorm.DB
}

func (this *UserModel) AuthenticateUser (username string, password string) (*User, error) {
	var user User

	// Find user db by param username
	var err = this.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Invalid credentials")
		}
		return nil, err
	}

	// Matches user db hashed password with param password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("Invalid credentials")
	}

	return &user, nil
}

func (this *UserModel) GetUserById(id string) (*User, error) {
	var user User

	var err = this.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
