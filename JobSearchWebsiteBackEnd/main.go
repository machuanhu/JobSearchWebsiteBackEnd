package main

import (
	"github.com/Job-Search-Website/pkg/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8080")
}
