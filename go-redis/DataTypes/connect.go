package main

import "github.com/redis/go-redis/v9"

type Pair struct {
	key   string
	value string
}

func NewRedisClient() (*redis.Client, error) {
	rClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rClient, nil
}
