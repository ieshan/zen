package conn

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type DBClient struct {
	*gorm.DB
}

type RedisClient struct {
	*redis.Client
}
