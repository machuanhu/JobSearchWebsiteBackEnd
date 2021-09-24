package dao
import (
	"github.com/Job-Search-Website/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) {
	db.Create(&user)
}
func IsEmailandPasswordMatched (email string, psw string) (isEmailandPasswordMatched bool) {
	var user models.User
	db.Where("email = ?", email).Find(&user)
	result := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(psw))
	if result ==nil {
		isEmailandPasswordMatched = true
	} else {
		isEmailandPasswordMatched = false
	}
	return
}
func EditIntroduction(email string,introduction string){
	db.Model(&models.User{}).Where("email = ?",email).Update("introduction",introduction)
}
func GetUser(email string)(user models.User,err bool){
	db.Where("email=?",email).Find(&user)
	return
}
func IsEmailRegistered(encrypted_email string) (IsRegistered bool) {
	var user models.User
	db.Where("email = ?", encrypted_email).Find(&user)
	if (user == models.User{}) {
		IsRegistered = false
	} else {
		IsRegistered = true
	}
	return IsRegistered
}