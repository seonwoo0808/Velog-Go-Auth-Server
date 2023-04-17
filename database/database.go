// pentag.kr/BuildinAuthVelog/database/database.go

package database

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// DB gorm connector
var DB *gorm.DB
var RDB *redis.Client
