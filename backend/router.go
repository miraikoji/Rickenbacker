package main

import (
	"net/http"
	"rickenbacker/controllers"
	"rickenbacker/models"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Router(e *echo.Echo, db *gorm.DB) {
	e.Use(CurrentUserLoader(db))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Rickenbacker!!")
	})

	sessionController := &controllers.SessionController{DB: db}
	e.POST("/login", sessionController.LoginHandler)
	e.GET("/secret_page", sessionController.SecretsPageHandler, UserAuthenticator)

	postController := &controllers.PostsController{DB: db}
	e.GET("/posts", postController.GetAllPosts)
	e.GET("/posts/:id", postController.GetPost)
	e.POST("/posts", postController.CreatePost, UserAuthenticator)
	e.PATCH("/posts/:id", postController.UpdatePost, UserAuthenticator)

	// userController := &controllers.UserController{DB: db}
}

func CurrentUserLoader(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := session.Get("session", c)
			if err != nil {
				return next(c)
			}

			if userID, ok := sess.Values["user_id"].(uint); ok {
				var user models.User
				if err := db.Take(&user, userID).Error; err == nil {
					c.Set("CurrentUser", &user)
				}
			}

			return next(c)
		}
	}
}

func UserAuthenticator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		if loggedIn, ok := sess.Values["logged_in"].(bool); !ok || !loggedIn {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		return next(c)
	}
}
