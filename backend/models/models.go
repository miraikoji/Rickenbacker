package models

import (
	"time"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Category{})
}

type BaseModel struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type User struct {
	BaseModel
	Email    string `form:"email"`
	Password string `form:"password"`
	Name     string `form:"name"`
	Posts    []Post
}

type Post struct {
	BaseModel
	Title      string `form:"title"`
	Body       string `form:"body"`
	UserID     uint
	CategoryID uint
}

type Category struct {
	BaseModel
	Title string `form:"title"`
	Posts []Post
}
