package models

import (
	"personal-blog-system/src/database"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	// Comments: 使用 Comment.PostAuthor 作为外键，Post 删除时级联删除评论
	Comments []Comment `gorm:"foreignKey:PostAuthor;constraint:OnDelete:CASCADE;" json:"comments,omitempty"`

	// Author 存储 post 所属用户的 ID
	Author uint `json:"author"`
	// 指定 User 的外键为 Author（即 Author -> users.id）
	User *User `gorm:"foreignKey:Author;references:ID" json:"user,omitempty"`
}

type Comment struct {
	gorm.Model
	Content    string `json:"content"`
	PostAuthor uint   `json:"post_author"`
	// 指定 Post 的外键为 PostAuthor
	Post *Post `gorm:"foreignKey:PostAuthor;references:ID" json:"post,omitempty"`

	// Commenter 存储评论者的用户 ID
	Commenter uint `json:"commenter"`
	// 指定 User 的外键为 Commenter（即 Commenter -> users.id）
	User *User `gorm:"foreignKey:Commenter;references:ID" json:"user,omitempty"`
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
