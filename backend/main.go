package main

import (
	"fmt"
	"net/http"
	"os"
	"rickenbacker/models"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "iamapen:password@tcp(db:3306)/rickenbacker?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました！", err)
		return
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

	Router(e, db)

	e.Logger.Fatal(e.Start(":9090"))
}
