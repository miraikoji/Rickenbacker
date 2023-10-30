package models

import (
	"time"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
}

type BaseModel struct {
	ID        uint      `param:"id"`
	CreatedAt time.Time `gorm:"<-create"`
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"<-delete"`
}

type User struct {
	BaseModel
	Email    string
	Password string
	Name     string
	Posts    []Post
}

type Post struct {
	BaseModel
	Title  string `form:"title"`
	Body   string `form:"body"`
	UserID uint   `form:"user_id"`
}
