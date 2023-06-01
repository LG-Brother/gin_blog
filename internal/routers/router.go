package routers

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	// v1版本路由组
	apiv1 := r.Group("/api/v1")
	{
		// 标签路由-----
		// 新增标签
		apiv1.POST("/tags")
		// 删除标签
		apiv1.DELETE("/tags/:id")
		// 修改标签
		apiv1.PUT("/tags/:id")
		// 部分修改标签：状态
		apiv1.PATCH("/tags/:id/state")
		// 获取所有标签
		apiv1.GET("/tags")
		// 文章路由-----
		// 新增文章
		apiv1.POST("/articles")
		// 删除文章
		apiv1.DELETE("/articles/:id")
		// 修改文章
		apiv1.PUT("/articles/:id")
		// 部分修改文章：状态
		apiv1.PATCH("/articles/:id/state")
		// 获取所有文章
		apiv1.GET("/articles")
		// 通过id获取文章
		apiv1.GET("/articles/:id")
	}
	return r
}
