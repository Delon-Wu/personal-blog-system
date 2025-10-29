package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string
	Content  string
	Comments []Comment
}

type Comment struct {
	gorm.Model
	Content string
	Post    Post `gorm:"foreignKey:PostID"`
	User    User `gorm:"foreignKey:UserID"`
}
