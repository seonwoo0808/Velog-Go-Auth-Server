// pentag.kr/BuildinAuthVelog/database/connect.go

package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"pentag.kr/BuildinAuthVelog/configs"
)

func ConnectDB() {
	var err error // define error here to prevent overshadowing the global DB

	dbConfig := configs.Config.DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Seoul", dbConfig.Host, dbConfig.User, dbConfig.Pass, dbConfig.Database, dbConfig.Port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// err = DB.AutoMigrate(
	// 	&models.User{},
	// )
	if err != nil {
		panic(err)
	}
}

func ConnectRedis() {
	redisConfig := configs.Config.Redis
	ctx := context.Background()
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       0, // use default DB
	})
	RDB.Ping(ctx)
}
