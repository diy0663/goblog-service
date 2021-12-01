package model

type Article struct {
	*Model
	Title         string `json:"title"`
	Description   string `json:"description"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (Article) TableName() string {
	return "blog_article"
}
