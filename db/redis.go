package db

import (
	"github.com/go-redis/redis"
	"github.com/ieshan/zen/config"
	"github.com/ieshan/zen/logger"
)

// Setup assigns redis.Client interface based on config to RedisClient
func ConnectRedis() {
	redisOnce.Do(func() {
		cfg := config.GetRedis()
		c := redis.NewClient(&redis.Options{
			Addr:     cfg.Address,
			Password: cfg.Password,
			DB:       cfg.DB,
		})

		log := logger.GetLogger()
		log.ErrorFatal("Got error when connecting with MySQL database", c.Ping().Err())

		redisConn = &RedisClient{}
		redisConn.Client = c
	})
}

// DefaultRedis returns the default RedisClient currently in Use
func GetRedisConn() *RedisClient {
	return redisConn
}

