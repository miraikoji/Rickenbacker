package controllers

import (
	"net/http"
	"rickenbacker/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PostsController struct {
	DB *gorm.DB
}

func (con *PostsController) GetAllPosts(c echo.Context) error {
	var posts []models.Post

	if result := con.DB.Find(&posts); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, posts)
}

func (con *PostsController) GetPost(c echo.Context) error {
	var post models.Post

	if result := con.DB.Take(&post, c.Param("id")); result.Error != nil {
		return c.JSON(http.StatusNotFound, "Post Not Found")
	}

	return c.JSON(http.StatusOK, post)
}
