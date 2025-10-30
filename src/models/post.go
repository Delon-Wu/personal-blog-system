package models

import (
	"personal-blog-system/src/database"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Comments []Comment `gorm:"foreignKey:PostAuthor;constraint:OnDelete:CASCADE;" json:"comments"`
	Author   uint      `json:"author"`
	User     *User     `gorm:"foreignKey:Author" json:"user,omitempty"`
}

type Comment struct {
	gorm.Model
	Content    string `json:"content"`
	PostAuthor uint   `json:"post_author"`
	Post       *Post  `gorm:"foreignKey:PostAuthor" json:"post,omitempty"`
	Commenter  uint   `json:"commenter"`
	User       *User  `gorm:"foreignKey:Commenter" json:"user,omitempty"`
}

func CreatePost(p *Post) (err error) {
	db := database.DB
	return db.Create(p).Error
}

func GetPostById(id uint) (post Post, err error) {
	db := database.DB
	err = db.Where("id = ?", id).First(&post).Error
	return post, err
}

func GetPostCommentById(id uint) (comments []Comment, err error) {
	db := database.DB
	// Comment uses PostAuthor as the foreign key column (snake_case -> post_author)
	err = db.Model(&Comment{}).Where("post_author = ?", id).Find(&comments).Error
	return comments, err
}

func GetPostList(page int, pageSize int) (posts []Post, err error) {
	db := database.DB
	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&posts).Error
	return posts, err
}

func DeletePost(id uint) (err error) {
	db := database.DB
	// With constraint:OnDelete:CASCADE on Comments, deleting the post will cascade in DB.
	// Use the model type delete; if you prefer to ensure comments removed at app level,
	// preload and delete associations explicitly.
	err = db.Where("id = ?", id).Delete(&Post{}).Error
	return err
}

func EditPost(p *Post, id uint) (err error) {
	db := database.DB
	return db.Model(&Post{}).Where("id = ?", id).Updates(p).Error
}

func CreateComment(c *Comment) (err error) {
	db := database.DB
	return db.Create(c).Error
}

func DeleteComment(id uint) (err error) {
	db := database.DB
	return db.Where("id = ?", id).Delete(&Comment{}).Error
}
