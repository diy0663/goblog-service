package setting

import "time"

// 定义对应yaml文件的配置项结构体

//  ServerSettingS 对应 configs/config.yaml 里面的 Server 区块
type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSettingS struct {
	DBType         string
	UserName       string
	Password       string
	Host           string
	Port           string
	DBName         string
	TablePrefix    string
	Charset        string
	ParseTime      bool
	MaxIdleConns   int
	MaxOpenConns   int
	MaxLifeSeconds int
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}
