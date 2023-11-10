package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"rickenbacker/models"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SessionController struct {
	DB *gorm.DB
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (sh *SessionController) LoginHandler(c echo.Context) error {
	req := LoginRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	email := req.Email
	password := req.Password

	var user models.User
	result := sh.DB.Where("Email = ?", email).First(&user)
	c.Logger().Debugf("user: %s", user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.String(http.StatusUnauthorized, "Invalid Email or Password")
		}
		// ErrRecordNotFound以外でエラーがでた場合
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return c.String(http.StatusUnauthorized, "Invalid Email or Password")
	}

	sess, _ := session.Get("session", c)
	sess.Values["logged_in"] = true
	sess.Values["user_id"] = user.ID
	sess.Save(c.Request(), c.Response())

	return c.String(http.StatusOK, "Logged in!")
}

func (sh *SessionController) SecretsPageHandler(c echo.Context) error {
	user, ok := c.Get("CurrentUser").(*models.User)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	message := fmt.Sprintf("CurrentUser userID: %v, userName: %v", user.ID, user.Name)
	return c.String(http.StatusOK, message)
}
