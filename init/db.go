package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"soleaf.xyz/yaowen/global"
	"time"
)



type MysqlInfo struct {
	User string
	Password string
	Host string
	Port int
	TableName string
}

func InitDB() {
	//dsn := "root:root@tcp(172.30.198.210:13306)/mxshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"

	cf := MysqlInfo{
		"root",
		"Mysql:12138",
		"219.228.135.107",
		3306,
		"yaowen",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cf.User, cf.Password, cf.Host, cf.Port, cf.TableName)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	var err error

	// 全局模式
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})

	if err != nil {
		fmt.Errorf("初始化gorm失败")
		panic(err)
	}
}
