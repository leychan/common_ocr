package redis

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)
var ctx = context.Background()

func newClient() (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,  // use default DB
	})
	_, err := db.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("连接错误, %s", err)
		return nil, err
	}
	return db, nil
}

func Get(key string) (string, error) {
	client, err := newClient()
	if err != nil {
		return "", err
	}
	return client.Get(ctx, key).Result()
}

func Set(key string, data interface{}, expire int64) (string, error) {
	client, err := newClient()
	if err != nil {
		return "", err
	}
	fmt.Println("expire: ", expire, ",   key: ", key)
	return client.Set(ctx, key, data, time.Duration(expire) * time.Second).Result()
}
