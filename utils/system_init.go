package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB     *gorm.DB
	Status map[string]string
	Red    *redis.Client
)

func InitConfig() {
	Status = map[string]string{}
	viper.SetConfigName("app") // name of config file (without extension)
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	Status = viper.GetStringMapString("status")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config status", Status["phonehasexisted"])
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

// 初始化redis
func InitRedis() {
	ip := viper.GetString("redis.ip")
	port := viper.GetString("redis.port")
	password := viper.GetString("redis.password")
	poolSize := viper.GetInt("redis.poolSize")
	db := viper.GetInt("mysql.db")
	minIdleConn := viper.GetInt("redis.minIdleConn")
	Red = redis.NewClient(&redis.Options{
		Addr:         ip + ":" + port,
		Password:     password,
		PoolSize:     poolSize,
		DB:           db,
		MinIdleConns: minIdleConn,
	})
	var ctx = context.Background()
	statusCmd := Red.Ping(ctx)
	fmt.Println("redis init ... ", statusCmd)
}

const (
	PublishKey = "websocket"
)

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	err := Red.Publish(ctx, channel, msg).Err()
	return err
}

// subscribe 订阅消息到Redis
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	fmt.Println("Subscribe .... ", sub)
	msg, err := sub.ReceiveMessage(ctx)
	return msg.Payload, err
}
