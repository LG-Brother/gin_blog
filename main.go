package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 设置Gin为发布模式
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.Run()
}
