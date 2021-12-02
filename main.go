package main

import (
	"log"
	"net/http"
	"time"

	"github.com/diy0663/goblog-service/global"
	"github.com/diy0663/goblog-service/internal/routers"
	"github.com/diy0663/goblog-service/pkg/setting"
)

// 自动初始化, 读取加载配置(Server,APP,database)
// 程序执行顺序 : 全局变量初始化 =>init 方法 => main 方法
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func main() {

	router := routers.NewRouter()
	// 由于配置已经在init加载了,所以在这里直接读取配置,而不是写死
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}

// 包内调用, 被init使用,用于初始化配置(读取配置)
func setupSetting() error {
	// 先new一个setting
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	// 对应setting进行操作
	// setting.ReadSection 里面已经指定了配置文件,从里面读取  Server 区块内容,并回写到一个全局配置中(这个全局配置的结构跟Server区块一致)
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	// 为配置里面的数字指定计量单位 秒
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}
