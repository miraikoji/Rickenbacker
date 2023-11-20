package models

import (
	"log"

	"github.com/go-faker/faker/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	seedUser := User{
		Email:    "mirai@example.com",
		Password: hashed("password"),
		Name:     "Taro Mirai",
	}
	db.FirstOrCreate(&seedUser)
	log.Println("[Seed] Create User:", seedUser.Email)

	for i := 0; i < 5; i++ {
		post := Post{
			Title:  faker.Sentence(),
			Body:   faker.Paragraph(),
			UserID: seedUser.ID,
		}
		db.FirstOrCreate(&post)
		log.Printf("[Seed] Create Post: %s", post.Title)
	}
}

func hashed(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hashedPassword)
}
