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

func (con *PostsController) CreatePost(c echo.Context) error {
	var post models.Post

	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	con.DB.Create(&post)
	return c.JSON(http.StatusCreated, post)
}

func (con *PostsController) UpdatePost(c echo.Context) error {
	var post models.Post

	if err := c.Bind(&post); err = nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	con.DB.Save(&post)
	return c.JSON(http.StatusOK, post)
}
