package models

import (
	"personal-blog-system/src/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Posts    []Post
	Comments []Comment
}

// HashPassword 加密密码
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword 验证密码
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// CreateUser creates a new user record. Returns an error on failure.
func CreateUser(user *User) error {
	db := database.DB
	if err := user.HashPassword(); err != nil {
		return err
	}

	return db.Create(user).Error
}

// GetUserByID returns a user by ID (string id). Preloads Posts and Comments.
func GetUserByID(id uint) (*User, error) {
	db := database.DB

	var user User
	if err := db.Preload("Posts").Preload("Comments").Omit("password").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
