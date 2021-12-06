package model

import (
	"fmt"
	"time"

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
	// fmt.Println(DatabaseSetting)

	// alter table `blog_tag` convert to character set utf8mb4 COLLATE utf8mb4_unicode_ci;
	config := mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:33066)/blog_service?charset=utf8&parseTime=True&loc=Local",
		// todo 配置项这里要确保返回值是字符串
		// DSN: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
		// 	DatabaseSetting.UserName,
		// 	DatabaseSetting.Password,
		// 	DatabaseSetting.Host,
		// 	DatabaseSetting.Port,
		// 	DatabaseSetting.DBName,
		// 	DatabaseSetting.Charset,
		// )
	})
	// fmt.Println(config)

	db, err = gorm.Open(config, &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Warn),
	})
	// if global.ServerSetting.RunMode == "debug" {
	// 	db, err = gorm.Open(config, &gorm.Config{
	// 		// 开启调试SQL,Logger 可用来指定和配置 GORM 的调试器，例如说命令行打印 SQL 语句
	// 		Logger: gormlogger.Default.LogMode(gormlogger.Warn),
	// 	})
	// } else {

	// }

	fmt.Println(db)
	if err != nil {
		return db, err
	}

	// 在这里注册自定义的相关回调
	// db.Callback().Create().Replace()
	// db.Callback().Update().Register()
	// db.Callback().Delete().Remove()

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(DatabaseSetting.MaxOpenConns)
	sqlDB.SetMaxIdleConns(DatabaseSetting.MaxIdleConns)
	// 每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(DatabaseSetting.MaxLifeSeconds) * time.Second)
	// 开启调试模式??

	// todo  MaxIdleConns 和 MaxOpenConns 配置还没应用上
	return db, nil
}

// todo gorm 2.0 的自定义回调, 以及上面的 NewDBEngine 中注册自定义的回调还没有写
