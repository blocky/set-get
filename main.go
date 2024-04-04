package main

import (
	"log"
	"net/http"
)

func main() {
	rdb, err := NewRedisClient()
	if err != nil {
		log.Fatalf("Error creating Redis client: %v", err)
	}

	http.HandleFunc("/set", handlerSetValue(rdb))
	http.HandleFunc("/get", handlerGetValue(rdb))

	_ = http.ListenAndServe(":8080", nil)
}
