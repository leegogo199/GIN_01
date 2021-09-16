package util

import (
	"GINVUE/Model"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
)

func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklxzvbnmQQERTYUOIOPASDFGHJKLZXCVBNM")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user Model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
