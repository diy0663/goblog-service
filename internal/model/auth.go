package model

import "gorm.io/gorm"

type Auth struct {
	*Model
	Appkey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {

	var auth Auth
	db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?", a.Appkey, a.AppSecret, 0)
	err := db.First(&auth).Error
	// 需要检查是否 真的查得到这个appkey以及app_secret
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, nil
	}
	return auth, err

}
