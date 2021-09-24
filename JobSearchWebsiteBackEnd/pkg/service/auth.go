package service

import (
	"github.com/Job-Search-Website/models"
	"github.com/Job-Search-Website/pkg/consts"
	"github.com/Job-Search-Website/pkg/dao"
	"github.com/Job-Search-Website/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Signup(context *gin.Context) {
	email := context.PostForm("email")
	password := context.PostForm("password")
	name := context.PostForm("name")
	role := context.PostForm("role")
	passwordhash := util.HashWithSalt(password)
	if email==""{
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "邮箱不能为空",
		})
		return
	}
	if dao.IsEmailRegistered(email) {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "该用户已经注册",
		})
		return
	}
	user := models.User{
		Email: email, Password: passwordhash, Username: name, RegisterTimestamp: util.GetTimeStamp(),
		Role: role}
	dao.CreateUser(user)
	context.JSON(http.StatusOK, gin.H{
		"email": user.Email,
		"msg":   "注册成功",
	})
	return
}

func Login (context *gin.Context) {
	password := context.PostForm("password")
	email := context.PostForm("email")
	token, _ := util.GenerateToken(email,password)

	if dao.IsEmailandPasswordMatched(email ,password) {
		// token
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":"登录成功",
			"token": token,
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户名密码不匹配",
		})
	}
	return
}
func Introduction(context *gin.Context){
	email,_:=util.GetEmailFromToken(context)
	introduction:=context.PostForm("introduction")
	dao.EditIntroduction(email,introduction)
	context.JSON(http.StatusOK,gin.H{
		"msg":"编辑成功",
	})
	return
}
func GetMyself(context *gin.Context) {
	email,_:=util.GetEmailFromToken(context)
	user, _ := dao.GetUser(email)
	context.JSON(http.StatusOK, gin.H{
		"email":        user.Email,
		"name":         user.Username,
		"role":         user.Role,
		"introduction": user.Introduction,
	})
	return
}
func GetMyResume(context *gin.Context){
	email,_:=util.GetEmailFromToken(context)
	resume:=dao.GetResumeByJobSeeker(email)
	results:=[]map[string]interface{}{}
	for _,temp:=range resume{
		if temp.IsReplied {
			user,_:=dao.GetUser(temp.HrEmail)
			results = append(results, map[string]interface{}{
				"hremail":       temp.HrEmail,
				"resumecontext": temp.ResumeContext,
				"reply":         temp.Reply,
				"releasetime":   temp.ReleseTime,
				"replytime":     temp.ReplyTime,
				"hrname":		 user.Username,
				"hrintroduction":user.Introduction,
			})
		}else{
			user,_:=dao.GetUser(temp.HrEmail)
			results = append(results, map[string]interface{}{
				"hremail":       temp.HrEmail,
				"resumecontext": temp.ResumeContext,
				"releasetime":   temp.ReleseTime,
				"hrname":		 user.Username,
				"hrintroduction":user.Introduction,
				"reply":"还没有收到回复",
			})
		}
	}
	context.JSON(http.StatusOK,gin.H{
		"resume":results,
	})
	return
}
func GetMyJobSeeker(context *gin.Context){
	email,_:=util.GetEmailFromToken(context)
	resume:=dao.GetResumeByHr(email)
	results:=[]map[string]interface{}{}
	for _,temp:=range resume{
		if temp.IsReplied {
			user,_:=dao.GetUser(temp.JobSeekerEmail)
			results = append(results, map[string]interface{}{
				"jobseekeremail":temp.JobSeekerEmail,
				"resumecontext": temp.ResumeContext,
				"reply":         temp.Reply,
				"releasetime":   temp.ReleseTime,
				"replytime":     temp.ReplyTime,
				"jobseekername": user.Username,
				"resume_id":temp.ID,
			})
		}else{
			user,_:=dao.GetUser(temp.HrEmail)
			results = append(results, map[string]interface{}{
				"jobseekeremail":temp.JobSeekerEmail,
				"resumecontext": temp.ResumeContext,
				"releasetime":   temp.ReleseTime,
				"jobseekername": user.Username,
				"reply":"还没有回复",
				"resume_id":temp.ID,
			})
		}
	}
	context.JSON(http.StatusOK,gin.H{
		"resume":results,
	})
	return
}
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString:= context.GetHeader("token")
		claim, err := util.GetClaimFromToken(tokenString)
		if err != nil {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未认证，请先登录",
			})
			return
		}

		//using assert is very dangerous
		tokenTimeStamp := claim.(jwt.MapClaims)["timeStamp"].(float64)
		time := util.GetTimeStamp() - int64(tokenTimeStamp)
		if time > consts.EXPIRE_TIME_TOKEN {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Token过期，请重新登录",
			})
		}


	}
}