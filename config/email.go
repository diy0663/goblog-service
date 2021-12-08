package config

import (
	"github.com/diy0663/goblog-service/pkg/config"
)

func init() {
	config.Add("email", config.StrMap{

		// todo vscode  已设置 小写转大写 cmd+M
		"host":     config.Env("EMAIL_HOST", ""),
		"port":     config.Env("EMAIL_PORT", 465),
		"username": config.Env("EMAIL_USERNAME", ""),
		"password": config.Env("EMAIL_PASSWORD", ""),
		"is_ssl":   config.Env("EMAIL_IS_SSL", true),
		"from":     config.Env("EMAIL_FROM", ""),
		"to":       config.Env("EMAIL_TO", ""),
	})
}
