package configs

import (
	"os"
	"strconv"
)

type WASConfig struct {
	Host string
	Port string
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Pass     string
	Database string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type MainConfig struct {
	WAS   WASConfig
	DB    DBConfig
	Redis RedisConfig
}

func getEnv_s(key string) string { // 환경변수를 가져오고 없으면 에러를 발생시키는 함수
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	panic("Environment variable " + key + " not found")
}

func getAllEnv() MainConfig { // 모든 환경변수를 가져오는 함수
	var err error

	dbPort, err := strconv.Atoi(getEnv_s("POSTGRES_PORT"))

	if err != nil {
		panic("Environment variable DB_PORT is not a number")
	}
	redisPort, err := strconv.Atoi(getEnv_s("REDIS_PORT"))
	if err != nil {
		panic("Environment variable REDIS_PORT is not a number")
	}

	return MainConfig{
		WAS: WASConfig{
			Host: getEnv_s("WAS_HOST"),
			Port: getEnv_s("WAS_PORT"),
		},
		DB: DBConfig{
			Host:     getEnv_s("POSTGRES_HOST"),
			Port:     dbPort,
			User:     getEnv_s("POSTGRES_USER"),
			Pass:     getEnv_s("POSTGRES_PASSWORD"),
			Database: getEnv_s("POSTGRES_DB"),
		},
		Redis: RedisConfig{
			Host:     getEnv_s("REDIS_HOST"),
			Port:     redisPort,
			Password: getEnv_s("REDIS_PASSWORD"),
			DB:       0,
		},
	}
}

var Config = getAllEnv() // 환경변수 객체
