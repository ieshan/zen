package db

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"sync"
)

// MySQLClient holds the MySQL client instance
type DBClient struct {
	*gorm.DB
}
// RedisClient holds the redis client instance
type RedisClient struct {
	*redis.Client
}

// conn is an instance *gorm.DB
var dbConn *DBClient
// Redis is an instance *redis.Client{}
var redisConn *RedisClient

var dbOnce sync.Once
var redisOnce sync.Once
