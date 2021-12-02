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
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

// 读取配置的方法,用了指针,所以返回值那里只需要error即可
func (s *Setting) ReadSection(k string, v interface{}) error {
	// UnmarshalKey 匹配某一个字段
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return nil
	}
	return nil

}
