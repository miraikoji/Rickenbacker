package models

import (
	"log"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	seedUser := User{
		Email:    "mirai@example.com",
		Password: "$2a$10$5VotD2mOBoRj2At0wG7bw.qSZgylGZydJoEP38fqQyiRphsqf8NLa",
		Name:     "sample",
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
