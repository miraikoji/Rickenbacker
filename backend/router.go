package main

import (
	"net/http"
	contollers "rickenbacker/controllers"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Router(e *echo.Echo, db *gorm.DB) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Rickenbacker!!")
	})

	sessionController := &contollers.SessionController{DB: db}
	e.POST("/login", sessionController.LoginHandler)
	e.GET("/secret_page", sessionController.SecretsPageHandler, checkLoginMiddleware)

	// userController := &contollers.UserController{DB: db}
}

func checkLoginMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		if loggedIn, ok := sess.Values["logged_in"].(bool); !ok || !loggedIn {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		return next(c)
	}
}
