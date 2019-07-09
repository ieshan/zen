package conn

import (
	"github.com/go-redis/redis"
	"github.com/ieshan/zen/config"
	"github.com/ieshan/zen/logger"
)

func GetRedisConn(cfg *config.RedisConfig) *RedisClient {
	c := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	log := logger.GetLogger()
	log.ErrorFatal("Got error when connecting with MySQL database", c.Ping().Err())

	return &RedisClient{Client: c}
}
