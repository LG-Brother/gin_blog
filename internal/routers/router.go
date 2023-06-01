package routers

import (
	v1 "gin_blog/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	tag := v1.Tag{}
	article := v1.Article{}
	// v1版本路由组
	apiv1 := r.Group("/api/v1")
	{
		// 标签路由-----
		// 新增标签
		apiv1.POST("/tags", tag.Create)
		// 删除标签
		apiv1.DELETE("/tags/:id", tag.Delete)
		// 修改标签
		apiv1.PUT("/tags/:id", tag.Update)
		// 部分修改标签：状态
		apiv1.PATCH("/tags/:id/state", tag.Update)
		// 获取所有标签
		apiv1.GET("/tags", tag.List)
		// 文章路由-----
		// 新增文章
		apiv1.POST("/articles", article.Create)
		// 删除文章
		apiv1.DELETE("/articles/:id", article.Delete)
		// 修改文章
		apiv1.PUT("/articles/:id", article.Update)
		// 部分修改文章：状态
		apiv1.PATCH("/articles/:id/state", article.Update)
		// 获取所有文章
		apiv1.GET("/articles", article.Get)
		// 通过id获取文章
		apiv1.GET("/articles/:id", article.List)
	}
	return r
}
