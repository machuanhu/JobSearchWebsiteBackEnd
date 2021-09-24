package service

import (
	"github.com/Job-Search-Website/models"
	"github.com/Job-Search-Website/pkg/dao"
	"github.com/Job-Search-Website/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ReleaseBlog(context *gin.Context){
	email,_:=util.GetEmailFromToken(context)
	blog_context:=context.PostForm("blog_context")
	releasetime:=time.Now()
	blog:=models.Blog{
		Email: email,ReleseTime: releasetime,BlogContext: blog_context,
	}
	dao.CreateBlog(blog)
	context.JSON(http.StatusOK,gin.H{
		"msg":"发布成功",
	})
}
func GetBlog(context *gin.Context){
	blogs:=dao.GetBlog()
	results:=[]map[string]interface{}{}
	for _,temp:=range blogs{
		user,_:=dao.GetUser(temp.Email)
		results=append(results, map[string]interface{}{
			"email":temp.Email,
			"name":user.Username,
			"blog_context":temp.BlogContext,
			"blog_id":temp.ID,
		})
	}
	context.JSON(http.StatusOK,gin.H{
		"msg":results,
	})
}
