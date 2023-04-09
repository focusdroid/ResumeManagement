package models

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = Init()

var RDB = InitRedisDB()

var Enforcer = InitCasbin()

func Init() *gorm.DB {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/resume?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	DB.Set("gorm:table_options", "ENGINE=InnDB").AutoMigrate(&User{}, &Resume{}, &Backlog{}, &BlackList{})
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err, "mysql数据库创建问题")
		}
	}()
	fmt.Println("mysql数据库连接成功!")
	return DB

}

func InitRedisDB() *redis.Client {
	var ctx = context.Background()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err, "redis服务未启动或连接不到redis")
		}
	}()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis服务未启动或连接不到redis")
		fmt.Println("redis-err", err)
	}
	fmt.Println("redis数据库连接成功！")
	return rdb
}

func InitCasbin() *casbin.Enforcer {
	a, err := xormadapter.NewAdapter("mysql", "root:mysql@tcp(127.0.0.1:3306)/resume?charset=utf8", true)
	if err != nil {
		log.Fatalf("casbin连接数据库错误error: adapter: %s", err)
	}

	e, err := casbin.NewEnforcer("models/conf/rbac_models.conf", a)
	if err != nil {
		log.Fatalf("初始化casbin错误error: model: %s", err)
	}
	fmt.Println("eeee", e)
	return e
}
