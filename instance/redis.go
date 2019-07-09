package instance

import (
	"github.com/ieshan/zen/config"
	"github.com/ieshan/zen/conn"
	"sync"
)

var redisOnce sync.Once
var redisConn *conn.RedisClient

func GetRedis() *conn.RedisClient {
	redisOnce.Do(func() {
		redisConn = conn.GetRedisConn(config.GetRedis())
	})
	return redisConn
}
