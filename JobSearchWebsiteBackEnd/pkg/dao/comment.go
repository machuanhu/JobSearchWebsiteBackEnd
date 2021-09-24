package dao

import "github.com/Job-Search-Website/models"

func CreateComment(comment models.Comment) {
	db.Create(&comment)
}
func GetComment(blog_id int)(comments []models.Comment){
	db.Where("blog_id=?",blog_id).Find(&comments)
	return comments
}
