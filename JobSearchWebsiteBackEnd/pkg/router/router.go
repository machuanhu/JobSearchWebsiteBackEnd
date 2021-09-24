package router

import (
	"github.com/Job-Search-Website/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "access-control-allow-origin", "token"},
		AllowCredentials: false,
	}
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	api := r.Group("/api")
	{
		user := api.Group("/auth")
		{
			user.Use(cors.Default())
			user.POST("/signup", service.Signup)
			user.POST("/login", service.Login)
			user.Use(service.Auth())
			user.POST("/introduction",service.Introduction)
			user.GET("/myself",service.GetMyself)
			user.GET("/myhr",service.GetMyResume)
			user.GET("/myjobseeker",service.GetMyJobSeeker)
		}
		resume := api.Group("/resume")
		{
			resume.Use(cors.Default())
			resume.Use(service.Auth())
			resume.POST("replyresume",service.ReplyResume)
			resume.POST("releaseresume",service.ReleaseResume)
		}
		hr := api.Group("/hr")
		{
			hr.Use(cors.Default())
			hr.GET("hr_list",service.GetHrList)
		}
		score := api.Group("/score")
		{
			score.Use(cors.Default())
			score.POST("resume_score",service.Score)
		}
		secret :=api.Group("/secret")
		{
			secret.Use(cors.Default())
			secret.POST("/resume_secret",service.Secret)
		}
		blog:=api.Group("/blog")
		{
			blog.Use(cors.Default())
			blog.GET("/blog_list",service.GetBlog)
			blog.GET("/comment_list",service.GetComment)
			blog.Use(service.Auth())
			blog.POST("/release_blog",service.ReleaseBlog)
			blog.POST("/release_comment",service.ReleaseComment)
		}
	}
}
