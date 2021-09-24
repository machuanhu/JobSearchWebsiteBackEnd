package dao

import "github.com/Job-Search-Website/models"


func CreateBlog(blog models.Blog) {
	db.Create(&blog)
}
func GetBlog()(blogs []models.Blog){
db.Find(&blogs)
return blogs
}

