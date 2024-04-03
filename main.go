package main

import (
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(
	address string,
	password string,
	database int,
) *redis.Client {
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     address,
			Password: password,
			DB:       database,
		})
	return rdb
}

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
		return
	}
	rdb := NewRedisClient(config.RedisAddress, config.RedisPassword, config.RedisDatabase)

	http.HandleFunc("/set", handlerSetValue(rdb))
	http.HandleFunc("/get", handlerGetValue(rdb))

	_ = http.ListenAndServe(":8080", nil)
}
