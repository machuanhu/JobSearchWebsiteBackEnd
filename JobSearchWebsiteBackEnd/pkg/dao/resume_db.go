package dao

import (
	"github.com/Job-Search-Website/models"
	"time"
)

func GetResumeByJobSeeker(email string)(resumes []models.Resume){
	db.Where("job_seeker_email=?",email).Find(&resumes)
	return
}
func GetResumeByHr(email string)(resumes []models.Resume){
	db.Where("hr_email=?",email).Find(&resumes)
	return
}
func ReplyResume(resume_id int,reply string,){
	replytime:=time.Now()
	db.Model(&models.Resume{}).Where("id = ?",resume_id).Update(models.Resume{Reply:reply,IsReplied: true,ReplyTime: replytime})
}
func CreateResume(resume models.Resume)  {
	db.Create(&resume)
}
