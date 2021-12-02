package setting

import "github.com/spf13/viper"

// 读取配置

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	// 设置读取路径
	vp.AddConfigPath("configs/")
	// 设置读取类型为yaml
	vp.SetConfigType("yaml")
	// 设置配置名称 ???
	vp.SetConfigName("config")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil

}
