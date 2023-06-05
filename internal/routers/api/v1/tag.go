package v1

import (
	"gin_blog/global"
	"gin_blog/internal/model"
	"gin_blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}
func (t Tag) List(c *gin.Context) {
	db := global.DBEngine
	var tags []model.Tag
	db.Find(&tags)
	c.JSON(errcode.Success.Code(), tags)
}

func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
