package redis

import (
	"po_go/conf"
	"po_go/utils"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func init() {
	logger := utils.Log()
	logger.Info("Redis init")
	var dbConfig = conf.Conf.Redisdb
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     dbConfig.Addr,
		Password: "",          // no password set
		DB:       dbConfig.Db, // use default DB
	})

	pong, err := RedisClient.Ping().Result()
	logger.Info("Redis ", pong, err)
}
