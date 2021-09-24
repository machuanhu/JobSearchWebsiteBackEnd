package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Job-Search-Website/models"
	"github.com/Job-Search-Website/pkg/consts"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func GetTimeStamp() (t int64) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t = time.Now().In(loc).Unix()
	return
}
func ReadSettingsFromFile(settingFilePath string) (config models.Config) {
	jsonFile, err := os.Open(settingFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Panic(err)
	}
	return config
}
func GetClaimFromToken(tokenString string) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(consts.TOKEN_SCRECT_KEY), err
	})
	if err != nil {
		return nil, err
	} else {
		claims = token.Claims.(jwt.MapClaims)
		return claims, nil
	}
}
func GetEmailFromToken(c *gin.Context) (email string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("您未登录，请登陆后查看")
		}
	}()
	tokenStr :=c.GetHeader("token")
	claim, _ := GetClaimFromToken(tokenStr)
	email = claim.(jwt.MapClaims)["email"].(string)
	return
}