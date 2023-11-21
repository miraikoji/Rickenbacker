package controllers

import (
	"net/http"
	"rickenbacker/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func (con *UserController) GetUser(c echo.Context) error {
	var user models.User

	if result := con.DB.Take(&user, c.Param("id")); result.Error != nil {
		return c.JSON(http.StatusNotFound, "User Not Found")
	}

	return c.JSON(http.StatusOK, user)
}

func (con *UserController) CreateUser(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if result := con.DB.Create(&user); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Couldn't create user")
	}
	return c.JSON(http.StatusCreated, user)
}
