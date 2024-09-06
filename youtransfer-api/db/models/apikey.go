package models

import (
	"gorm.io/gorm"
	"math/rand"
	"youtransfer/api/db"
)

type ApiKey struct {
	*gorm.Model
	Key string `gorm:"unique;not null"`
}

func GenerateApiKey() string {
	key := generateRandomString(32)
	if err := db.GetConnection().Create(&ApiKey{Key: key}).Error; err != nil {
		return GenerateApiKey()
	}
	return key
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}
