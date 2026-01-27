package main

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type StringRepo struct {
	c *redis.Client
}

func NewStringRepo() (*StringRepo, error) {
	client, err := NewRedisClient()
	if err != nil {
		return nil, err
	}

	return &StringRepo{
		c: client,
	}, nil
}

func (s *StringRepo) StoreString(p Pair) error {
	err := s.c.Set(context.Background(), p.key, p.value, time.Duration(time.Minute*3)).Err()
	return err
}

func (s *StringRepo) GetString(key string) (string, error) {
	name, err := s.c.Get(context.Background(), key).Result()
	return name, err
}

func (s *StringRepo) DeleteString(key string) error {
	err := s.c.Del(context.Background(), key).Err()
	return err
}

func (s *StringRepo) StoreMultipleString(pairs []Pair) error {
	if len(pairs) == 0 {
		return nil
	}

	values := make([]any, 0, len(pairs)*2)
	for _, p := range pairs {
		values = append(values, p.key, p.value)
	}

	return s.c.MSet(context.Background(), values...).Err()
}

func (s *StringRepo) GetMultipleString(keys ...string) ([]Pair, error) {
	res, err := s.c.MGet(context.Background(), keys...).Result()
	if err != nil {
		return nil, err
	}

	pairs := make([]Pair, len(keys))
	for i, key := range keys {
		var val string
		if res[i] != nil {
			val = res[i].(string)
		}

		pairs = append(pairs, Pair{
			key:   key,
			value: val,
		})
	}

	return pairs, nil
}

func (s *StringRepo) GetStringRange(key string, start int64, end int64) (string, error) {
	res, err := s.c.GetRange(context.Background(), key, start, end).Result()
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *StringRepo) InsertString(key string, offset int64, value string) error {
	err := s.c.SetRange(context.Background(), key, offset, value).Err()
	return err
}

func (s *StringRepo) IncrementVal(key string) error {
	err := s.c.Incr(context.Background(), key).Err()
	return err
}

func (s *StringRepo) DecrementVal(key string) error {
	err := s.c.Decr(context.Background(), key).Err()
	return err
}

func (s *StringRepo) IncrementValBy(key string, val int64) error {
	err := s.c.IncrBy(context.Background(), key, val).Err()
	return err
}

func (s *StringRepo) DecrementValBy(key string, val int64) error {
	err := s.c.DecrBy(context.Background(), key, val).Err()
	return err
}
