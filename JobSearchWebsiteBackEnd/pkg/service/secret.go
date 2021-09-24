package service

import (
	"github.com/Job-Search-Website/pkg/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Secret(context *gin.Context){
	resume:=context.PostForm("resume_context")
	score:=dao.Secret(resume)
	context.JSON(http.StatusOK,gin.H{
		"result":score,
	})
}
