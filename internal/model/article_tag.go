package model

type ArticleTag struct {
	*Model
	TagID     uint64 `json:"tag_id"`
	ArticleID uint64 `json:"article_id"`
}

func (ArticleTag) TableName() string {
	return "blog_article_tag"
}
