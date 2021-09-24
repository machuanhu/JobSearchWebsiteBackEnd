package util

import "golang.org/x/crypto/bcrypt"

func HashWithSalt(plainText string) (HashText string) {

	hash, _ := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.MinCost)
	HashText = string(hash)
	return
}
