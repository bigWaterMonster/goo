package utils

import (
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"

	"g1/auth/model"
)

func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyz")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Debug().Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
