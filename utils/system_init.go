package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app") // name of config file (without extension)
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config mysql", viper.Get("mysql"))
}

func InitMysql() {
	name := viper.GetString("mysql.name")
	password := viper.GetString("mysql.password")
	ip := viper.GetString("mysql.ip")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	dsn := name + ":" + password + "@tcp(" + ip + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(dsn)

	// 自定义日志文件
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 彩色
		},
	)
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})

}
