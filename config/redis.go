package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// RedisCfg holds the redis configuration
type RedisConfig struct {
	Address  string
	Password string
	DB       int
	Prefix   string
	TTL      time.Duration
}

var redisConfig *RedisConfig

// Redis returns the default Redis configuration
func GetRedis() *RedisConfig {
	return redisConfig
}

func initRedis() {
	if redisConfig != nil {
		return
	}
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)
	viper.SetDefault("redis.prefix", "")
	viper.SetDefault("redis.ttl", 3600)
	redisConfig = &RedisConfig{
		Address:  fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		Prefix:   viper.GetString("redis.prefix") + "_",
		TTL:      viper.GetDuration("redis.ttl") * time.Second,
	}
}
