package dao

import (
	"database/sql"
	"fmt"
	"github.com/Job-Search-Website/models"
	"github.com/Job-Search-Website/pkg/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
)
func dsn(settings models.DbSettings) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4,utf8", settings.Username, settings.Password, settings.Hostname, settings.Dbname)
}
var db *gorm.DB
func init(){
databaseInit()
}

func databaseInit() {
	conf := util.ReadSettingsFromFile("Config.json")
	settings := conf.DbSettings
	connStr := dsn(settings)

	dbStr := strings.Replace(connStr, settings.Dbname, "", 1)
	msdb, e := sql.Open("mysql", dbStr)
	util.CheckError(e)
	msdb.Exec("create database if not exists " + settings.Dbname + " character set utf8")
	msdb.Close()

	var err1 error
	db, err1 = gorm.Open("mysql", connStr)
	//db.DB().SetMaxIdleConns(0)
	util.CheckError(err1)

	var temp []interface{}
	var user models.User
	var resume models.Resume
	var blog models.Blog
	var comment	models.Comment
	temp = append(temp,user,resume ,blog,comment)
	util.CreateTableIfNotExist(db, temp)
}