package db

import (
	"fmt"
	"log"

	"github.com/oxygen-org/server/config"
	"github.com/go-redis/redis"
)

var RedisC *redis.Client

func init() {
	host := config.CONFIG.Get("REDIS_HOST").MustString()
	port := config.CONFIG.Get("REDIS_PORT").MustInt()
	addr := fmt.Sprintf("%s:%d",host,port)
	password := config.CONFIG.Get("REDIS_PASSWORD").MustString()
	db := config.CONFIG.Get("REDIS_DB").MustInt()
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, 
		DB:       db,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(pong, err)
	RedisC = client
}
