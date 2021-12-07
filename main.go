package main

import (
	"log"
	"net/http"
	"time"

	"github.com/diy0663/goblog-service/config"
	c "github.com/diy0663/goblog-service/pkg/config"

	"github.com/diy0663/goblog-service/global"
	"github.com/diy0663/goblog-service/internal/model"
	"github.com/diy0663/goblog-service/internal/routers"
	"github.com/diy0663/goblog-service/pkg/logger"
	"github.com/diy0663/goblog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 自动初始化, 读取加载配置(Server,APP,database)
// 程序执行顺序 : 全局变量初始化 =>init 方法 => main 方法
func init() {

	// 读取全局配置

	// 从.env中读取全局变量,// 把读取到的设置到全局变量里面去
	config.Initialize()
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	// // todo 初始化logger
	err = setUpLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	// todo 在这里就有问题得到一个经过配置了的gorm全局配置
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDB err: %v", err)
	}

}

// @title 博客系统
// @version 1.0
// @description  gin-api-study
// @termsOfService gin-api-study
func main() {

	router := routers.NewRouter()

	// 由于配置已经在init加载了,所以在这里直接读取配置,而不是写死
	s := &http.Server{
		Addr: ":" + c.GetString("server.http_port"),
		// Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    time.Duration(c.GetInt64("server.read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(c.GetInt64("server.write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}

//包内调用, 被init使用,用于初始化配置(读取配置)
func setupSetting() error {
	// 先new一个setting
	// setting, err := setting.NewSetting()
	// if err != nil {
	// 	return err
	// }

	// todo ,把.env里面的配置 设置到 global.ServerSetting里面去
	// todo 后续要验证一下是否能直接把一整个env配置项赋值进去
	ServerSetting := setting.ServerSettingS{
		RunMode:      c.GetString("server.run_mode"),
		HttpPort:     c.GetString("server.http_port"),
		ReadTimeout:  time.Duration(c.GetInt64("server.read_timeout")),
		WriteTimeout: time.Duration(c.GetInt64("server.write_timeout")),
	}

	global.ServerSetting = &ServerSetting

	AppSetting := setting.AppSettingS{
		DefaultPageSize: c.GetInt("app.default_page_size"),
		MaxPageSize:     c.GetInt("app.max_page_size"),
		LogSavePath:     c.GetString("app.log_save_path"),
		LogFileName:     c.GetString("app.log_file_name"),
		LogFileExt:      c.GetString("app.log_file_ext"),
	}
	global.AppSetting = &AppSetting

	// 全局的jwt配置读取
	JwtSetting := setting.JWTSettingS{
		Secret: c.GetString("jwt.secret"),
		Issuer: c.GetString("jwt.issuer"),
		Expire: time.Duration(c.GetInt("jwt.expire")),
	}
	global.JwtSetting = &JwtSetting
	//todo 时间单位转换
	global.JwtSetting.Expire *= time.Second

	DatabaseSetting := setting.DatabaseSettingS{
		DBType:         "mysql",
		UserName:       c.GetString("database.mysql.username"),
		Password:       c.GetString("database.mysql.password"),
		Host:           c.GetString("database.mysql.host"),
		Port:           c.GetString("database.mysql.port"),
		DBName:         c.GetString("database.mysql.database"),
		TablePrefix:    c.GetString("database.mysql.table_prefix"),
		Charset:        c.GetString("database.mysql.charset"),
		ParseTime:      true,
		MaxIdleConns:   c.GetInt("database.mysql.max_idle_connections"),
		MaxOpenConns:   c.GetInt("database.mysql.max_open_connections"),
		MaxLifeSeconds: c.GetInt("database.mysql.max_life_seconds"),
	}
	global.DatabaseSetting = &DatabaseSetting

	// 为配置里面的数字指定计量单位 秒
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	// global.DBEngine 本身是个指针,执行成功之后会加载gorm对象到里面去,而且他是全局变量
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil

}

func setUpLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt

	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	// 这里其实就有问题了!! global.Logger有问题
	// fmt.Println(global.Logger)

	return nil

}
