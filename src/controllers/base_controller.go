package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (bc *BaseController) Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"message": "success",
		"data":    data,
		"error":   nil,
	})
}

func (bc *BaseController) Error(c *gin.Context, message string) {
	c.JSON(200, gin.H{
		"message": "fail",
		"data":    nil,
		"error":   message,
	})
}

func (bc *BaseController) GetIDFromParam(c *gin.Context) (id uint, err error) {
	idStr := c.Param("id")
	id64, _err := strconv.ParseUint(idStr, 10, 64)
	if _err != nil {
		return 0, _err
	}
	return uint(id64), nil
}

func (bc *BaseController) Pagination(c *gin.Context) (page int, pageSize int) {
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ = strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	return page, pageSize
}
