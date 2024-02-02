package controllers

import (
	"net/http"
	"rickenbacker/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CategoryController struct {
	DB *gorm.DB
}

func (con *CategoryController) GetAllCategories(c echo.Context) error {
	var categories []models.Category

	if result := con.DB.Find(&categories); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, categories)
}

func (con *CategoryController) CreateCategory(c echo.Context) error {
	var category models.Category

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if result := con.DB.Create(&category); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Couldn't create category")
	}

	return c.JSON(http.StatusCreated, category)
}

func (con *CategoryController) UpdateCategory(c echo.Context) error {
	var category models.Category

	if result := con.DB.Take(&category, c.Param("id")); result.Error != nil {
		return c.JSON(http.StatusNotFound, "Category Not Found")
	}

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if result := con.DB.Save(&category); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Couldn't update category")
	}
	return c.JSON(http.StatusOK, category)
}
