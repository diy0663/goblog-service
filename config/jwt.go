package config

import "github.com/diy0663/goblog-service/pkg/config"

func init() {
	config.Add("jwt", config.StrMap{
		// jwt的整体密钥
		"secret": config.Env("JWT_SECRET", ""),
		// 签发者
		"issuer": config.Env("JWT_ISSUER", ""),
		// 过期时效
		"expire": config.Env("JWT_EXPIRE", 7200),
	})
}
