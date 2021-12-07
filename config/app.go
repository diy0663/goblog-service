package config

import "github.com/diy0663/goblog-service/pkg/config"

func init() {
	config.Add("app", config.StrMap{
		"default_page_size":       config.Env("DEFAULT_PAGE_SIZE", 10),
		"max_page_size":           config.Env("MAX_PAGE_SIZE", 100),
		"log_save_path":           config.Env("LOG_SAVE_PATH", ""),
		"log_file_name":           config.Env("LOG_FILE_NAME", ""),
		"log_file_ext":            config.Env("LOG_FILE_EXT", ""),
		"dfault_context_time_out": config.Env("DEFAULT_CONTEXT_TIME_OUT", 60),
	})
}
