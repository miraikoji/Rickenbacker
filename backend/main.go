package main

import (
	"fmt"
	"net/http"
	"os"
	"rickenbacker/models"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	maxRetry := 5
	retryInterval := time.Second * 5
	var db *gorm.DB
	var err error

	for i := 0; i < maxRetry; i++ {
		dsn := "iamapen:password@tcp(db:3306)/rickenbacker?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err == nil {
			fmt.Println("データベースに接続しました。")
			break
		}

		fmt.Println("データベースへの接続に失敗しました。リトライします…", err)
		time.Sleep(retryInterval)
	}

	models.Migrate(db)
	if os.Getenv("SEED_DATA") == "true" {
		models.Seed(db)
	}

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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	Router(e, db)

	e.Logger.Fatal(e.Start(":9090"))
}
