package util

import (
	"fmt"
	"github.com/Job-Search-Website/pkg/consts"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"log"
	"reflect"
)

func CheckError(err error) bool {
	if err != nil {
		log.Fatal(err.Error())
		return true
	}
	return false
}

func CreateTableIfNotExist(db *gorm.DB, tableModels []interface{}) {
	for _, value := range tableModels {
		if !db.HasTable(value) {
			db.CreateTable(value)
			fmt.Println("Create table ", reflect.TypeOf(value), " successfully")
		}
	}
}
func GenerateToken(email string, role string) (tokenString string, err error) {
	claims := jwt.MapClaims{
		"email":     email,
		"timeStamp": GetTimeStamp(),
		"role":      role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(consts.TOKEN_SCRECT_KEY))
	return
}
