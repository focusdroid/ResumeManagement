package models

import (
	"fmt"
	"github.com/go-redis/redis/v9"
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
	DB.AutoMigrate(&User{}, &Resume{}, &Backlog{})
	//DB.Migrator().AddColumn(&User{}, "IsAdmin")
	DB.Migrator().DropColumn(&User{}, "IsAdmins")
	//DB.Migrator().DropColumn(&User{}, "IsAdmin")
	fmt.Println("数据库连接成功!")
	return DB

}

func InitRedisDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
