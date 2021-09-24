package service

import (
	"github.com/Job-Search-Website/pkg/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetHrList(context *gin.Context){
	page,_:=strconv.Atoi(context.Query("page"))
	number,_:=strconv.Atoi(context.Query("number"))
	hrs:=dao.GetHrList(page,number)
	results:=[]map[string]interface{}{}
	for _,temp:=range hrs{
		results=append(results, map[string]interface{}{
			"email":temp.Email,
			"name":temp.Username,
			"introduction":temp.Introduction,
		})
	}
	context.JSON(http.StatusOK,gin.H{
		"msg":results,
	})
}
