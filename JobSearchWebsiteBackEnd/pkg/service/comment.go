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

func ReleaseComment(context *gin.Context){
	email,_:=util.GetEmailFromToken(context)
	comment_context:=context.PostForm("comment_context")
	blog_id,_:=strconv.Atoi(context.PostForm("blog_id"))
	releasetime:=time.Now()
	comment:=models.Comment{
		Email: email,ReleseTime: releasetime,CommentContext: comment_context,BlogId: blog_id,
	}
	dao.CreateComment(comment)
	context.JSON(http.StatusOK,gin.H{
		"msg":"发布成功",
	})
}

func GetComment(context *gin.Context){
	blog_id,_:=strconv.Atoi(context.Query("blog_id"))
	comments:=dao.GetComment(blog_id)
	results:=[]map[string]interface{}{}
	for _,temp:=range comments{
		user,_:=dao.GetUser(temp.Email)
		results=append(results, map[string]interface{}{
			"comment_context":temp.CommentContext,
			"release_time":temp.ReleseTime,
			"name":user.Username,
		})
	}
	context.JSON(http.StatusOK,gin.H{
		"msg":results,
	})
}
