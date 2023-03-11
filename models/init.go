package models

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = Init()

var RDB = InitRedisDB()

func Init() *gorm.DB {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/resume?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	DB.AutoMigrate(&User{}, &Resume{}, &Backlog{}, &BlackList{})
	fmt.Println("mysql数据库连接成功!")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err, "mysql数据库创建问题")
		}
	}()
	return DB

}

var ctx = context.Background()

func InitRedisDB() *redis.Client {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err, "redis可能是数据库创建问题")
		}
	}()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis服务未启动或连接不到redis")
		fmt.Println("redis-err", err)
	}
	return rdb
}
