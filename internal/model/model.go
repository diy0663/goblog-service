package model

import (
	"fmt"

	"github.com/diy0663/goblog-service/global"
	"github.com/diy0663/goblog-service/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// 创建模型基类 把共用字段放这里
type Model struct {
	ID         uint64 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	CreatedOn  uint64 `json:"created_on"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn uint64 `json:"modified_on"`
	DeletedOn  uint64 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

//把创建gorm 的初始化实例方法放这里,外层把放数据库全局配置的结构体
func NewDBEngine(DatabaseSetting *setting.DatabaseSettingS) (db *gorm.DB, err error) {
	config := mysql.New(mysql.Config{
		//DSN: "root:123456@tcp(127.0.0.1:33066)/goblog?charset=utf8&parseTime=True&loc=Local",
		DSN: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
			DatabaseSetting.UserName,
			DatabaseSetting.Password,
			DatabaseSetting.Host,
			DatabaseSetting.Port,
			DatabaseSetting.DBName,
			DatabaseSetting.Charset,
		)})

	// 根据配置选择开启debug模式
	if global.ServerSetting.RunMode == "debug" {
		db, err = gorm.Open(config, &gorm.Config{
			// 开启调试SQL,Logger 可用来指定和配置 GORM 的调试器，例如说命令行打印 SQL 语句
			Logger: gormlogger.Default.LogMode(gormlogger.Warn),
		})
	} else {
		db, err = gorm.Open(config, &gorm.Config{})
	}

	if err != nil {
		return db, err
	}

	// todo  MaxIdleConns 和 MaxOpenConns 配置还没应用上
	return db, nil
}
