package model

import "github.com/diy0663/goblog-service/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (Tag) TableName() string {
	return "blog_tag"
}

// tag.go 为了让swagger注解中的返回数据更贴合实际
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
