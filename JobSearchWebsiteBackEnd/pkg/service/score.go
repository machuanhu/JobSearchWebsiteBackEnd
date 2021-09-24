package service

import (
	"github.com/Job-Search-Website/pkg/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Score(context *gin.Context){
	resume:=context.PostForm("resume")
	score:=dao.Score(resume)
	context.JSON(http.StatusOK,gin.H{
		"result":score,
	})
}
