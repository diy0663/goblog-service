package dao

import "github.com/diy0663/goblog-service/internal/model"

// dao 里面去调用对应的 model方法

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{
		Appkey:    appKey,
		AppSecret: appSecret,
	}
	return auth.Get(d.engine)
}
