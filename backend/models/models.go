package models

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
}

type User struct {
	gorm.Model
	Email    string
	Password string
	Name     string
	Posts    []Post
}

type Post struct {
	gorm.Model
	Title  string
	Body   string
	UserID uint
}
