package handler

import (
	"github.com/go-redis/redis/v7"
)

func Redis() (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = client.Ping().Result()
	if err != nil {
		return
	}
	return client, err
	// Output: PONG <nil>
}
