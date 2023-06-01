package model

type ArticleTag struct {
	ID        uint32 `gorm:"primary_key" json:"id"`
	TagID     uint32 `json:"tag_id"`
	ArticleId uint32 `json:"article_id"`
}

func (a *ArticleTag) TableName() string {
	return "blog_article_tag"
}
