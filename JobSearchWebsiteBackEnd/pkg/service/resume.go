package service

import (
	"github.com/Job-Search-Website/models"
	"github.com/Job-Search-Website/pkg/dao"
	"github.com/Job-Search-Website/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func ReplyResume(context *gin.Context){
	resume_id,_:=strconv.Atoi(context.PostForm("resume_id"))
	reply:=context.PostForm("reply")
	dao.ReplyResume(resume_id,reply)
	context.JSON(http.StatusOK,gin.H{
		"msg":"回复成功",
	})
}
func ReleaseResume(context *gin.Context){
	job_seeker_email,_:=util.GetEmailFromToken(context)
	hr_email:=context.PostForm("hr_email")
	resume_context:=context.PostForm("resume_context")
	releasetime:=time.Now()
	resume:=models.Resume{
		JobSeekerEmail:job_seeker_email,HrEmail:hr_email,IsReplied:false,ResumeContext:resume_context,ReleseTime: releasetime,ReplyTime:releasetime,
	}
	dao.CreateResume(resume)
	context.JSON(http.StatusOK,gin.H{
		"msg":"发布成功",
	})
	return
}
