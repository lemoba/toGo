package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
)

// RedisClient
var (
	RedisClient *redis.Client
	RedisAddr   string
	RedisPwd    string
	RedisDB     int
)

func init() {
	cfg, err := ini.Load("../../.ini")
	if err != nil {
		panic(err)
	}
	LoadRedis(cfg)
	RedisInit()
}

func RedisInit() {
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPwd,
		DB:       RedisDB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}

func LoadRedis(cfg *ini.File) {
	host := cfg.Section("redis").Key("HOST").String()
	port := cfg.Section("redis").Key("PORT").String()
	RedisAddr = fmt.Sprintf("%s:%s", host, port)
	RedisPwd = cfg.Section("redis").Key("PASSWORD").String()
	RedisDB = cfg.Section("redis").Key("DB").MustInt()
}
