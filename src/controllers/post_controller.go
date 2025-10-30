package controllers

import (
	"personal-blog-system/src/models"

	"github.com/gin-gonic/gin"
)

var Post = PostController{}

type PostController struct {
	BaseController
}

type PostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  uint   `json:"user_id"`
}

func (pc *PostController) Create(c *gin.Context) {
	req := &PostRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		pc.Error(c, err.Error())
		return
	}

	userID, _ := Auth.GetUserID(c)
	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		Author:  userID,
	}

	if err := models.CreatePost(&post); err != nil {
		pc.Error(c, err.Error())
		return
	}

	// 返回创建的帖子（可以改为 DTO）
	pc.Success(c, post)
}

func (pc *PostController) Get(c *gin.Context) {
	id, _ := pc.GetIDFromParam(c)
	post, err := models.GetPostById(id)
	if err != nil {
		pc.Error(c, err.Error())
		return
	}
	pc.Success(c, post)
}

func (pc *PostController) List(c *gin.Context) {
	page, pageSize := pc.Pagination(c)
	posts, err := models.GetPostList(page, pageSize)
	if err != nil {
		pc.Error(c, err.Error())
	}
	pc.Success(c, posts)
}

func (pc *PostController) Edit(c *gin.Context) {
	req := struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Id      uint   `json:"id" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		pc.Error(c, err.Error())
	}
	userID, _ := Auth.GetUserID(c)
	post, _ := models.GetPostById(req.Id)
	if userID != post.Author {
		pc.Error(c, "It's not the owner of the post")
		return
	}

	if err := models.EditPost(&models.Post{
		Title:   req.Title,
		Content: req.Content,
	}, req.Id); err != nil {
		pc.Error(c, err.Error())
		return
	}

	pc.Success(c, post)
}

func (pc *PostController) Delete(c *gin.Context) {
	id, _ := pc.GetIDFromParam(c)

	// 获取帖子并检查是否存在
	post, err := models.GetPostById(id)
	if err != nil {
		pc.Error(c, err.Error())
		return
	}

	// 获取当前用户 ID 并验证是否为作者
	userID, ok := Auth.GetUserID(c)
	if !ok {
		pc.Error(c, "unauthenticated")
		return
	}

	if post.Author != userID {
		pc.Error(c, "you are not the owner of this post")
		return
	}

	if err := models.DeletePost(id); err != nil {
		pc.Error(c, err.Error())
		return
	}
	pc.Success(c, nil)
}

func (pc *PostController) GetComments(c *gin.Context) {
	id, _ := pc.GetIDFromParam(c)
	post, err := models.GetPostCommentById(id)
	if err != nil {
		pc.Error(c, err.Error())
		return
	}
	pc.Success(c, post)
}

func (pc *PostController) CreateComment(c *gin.Context) {
	req := struct {
		Content string `json:"content" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil || req.Content == "" {
		pc.Error(c, "Content is required")
		return
	}
	userID, ok := Auth.GetUserID(c)
	postID, e2 := pc.GetIDFromParam(c)
	if !ok || e2 != nil {
		pc.Error(c, "You are not the owner of this post")
		return
	}
	err := models.CreateComment(&models.Comment{
		Content:    req.Content,
		Commenter:  userID,
		PostAuthor: postID,
	})
	if err != nil {
		pc.Error(c, err.Error())
		return
	}
	pc.Success(c, nil)
}

func (pc *PostController) DeleteComment(c *gin.Context) {
	id, err := pc.GetIDFromParam(c)
	if err != nil {
		pc.Error(c, err.Error())
		return
	}
	err = models.DeleteComment(id)
	if err != nil {
		pc.Error(c, err.Error())
		return
	}
	pc.Success(c, nil)
}
