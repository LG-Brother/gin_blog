package main

import (
	"gin_blog/internal/routers"
	"net/http"
	"time"
)

func main() {
	// 设置Gin为发布模式
	//gin.SetMode(gin.ReleaseMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
