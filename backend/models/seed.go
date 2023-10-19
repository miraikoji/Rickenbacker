package models

import (
	"log"

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
}
