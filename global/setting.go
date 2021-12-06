package global

import (
	"github.com/diy0663/goblog-service/pkg/logger"
	"github.com/diy0663/goblog-service/pkg/setting"
)

// 全局配置

var (
	// http服务配置(端口之类的)
	ServerSetting *setting.ServerSettingS
	// app项目配置(分页之类的配置信息)
	AppSetting *setting.AppSettingS
	// 数据库之类的配置
	DatabaseSetting *setting.DatabaseSettingS

	JwtSetting *setting.JWTSettingS

	Logger *logger.Logger
)
