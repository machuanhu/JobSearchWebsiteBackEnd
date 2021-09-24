package dao

import "github.com/Job-Search-Website/models"

func GetHrList(page int,number int)(Hrs []models.User){
	db.Offset(page*number).Limit(number).Where("role=?","hr").Find(&Hrs)
	return Hrs
}
