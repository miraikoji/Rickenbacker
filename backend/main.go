package main

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId   string
	Password string
}

type LoginRequest struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

// 動作確認用のデータ
// 実際のアプリケーションでは、データベースを使う
var users = map[string]string{
	"t4traw": "$2a$10$5VotD2mOBoRj2At0wG7bw.qSZgylGZydJoEP38fqQyiRphsqf8NLa",
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	store := sessions.NewCookieStore([]byte("secret-key"))
	store.Options = &sessions.Options{
		Path: "/",
		// MaxAge:   86400,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	e.Use(session.Middleware(store))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Rickenbacker!!")
	})

	e.GET("/login", loginPageHandler)
	e.POST("/login", loginHandler)
	e.GET("/secure_page", securePageHandler, checkLoginMiddleware)

	e.Logger.Fatal(e.Start(":9090"))
}

func loginPageHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Please login")
}

func loginHandler(c echo.Context) error {
	req := LoginRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	userId := req.UserId
	password := req.Password

	hashedPassword, exists := users[userId]
	if !exists {
		return c.String(http.StatusUnauthorized, "Invalid userId or Password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return c.String(http.StatusUnauthorized, "Invalid userId or Password")
	}

	sess, _ := session.Get("session", c)
	sess.Values["logged_in"] = true
	sess.Values["user_id"] = userId
	sess.Save(c.Request(), c.Response())

	return c.String(http.StatusOK, "Logged in!")
}

func securePageHandler(c echo.Context) error {
	return c.String(http.StatusOK, "secure page")
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
