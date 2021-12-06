package config

import "github.com/diy0663/goblog-service/pkg/config"

func init() {
	config.Add("server", config.StrMap{
		"run_mode":      config.Env("RUN_MODE", "debug"),
		"http_port":     config.Env("HTTP_PORT", 8080),
		"read_timeout":  config.Env("READ_TIMEOUT", 60),
		"write_timeout": config.Env("WRITE_TIMEOUT", 60),
	})
}
