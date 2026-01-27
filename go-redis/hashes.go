package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type HashesRepo struct {
	c *redis.Client
}

func HashRepo() (*HashesRepo, error) {
	client, err := NewRedisClient()
	if err != nil {
		return nil, err
	}

	return &HashesRepo{
		c: client,
	}, nil
}

func (h *HashesRepo) Set(key string, value []Pair) error {
	valMap := make(map[string]string)
	for _, val := range value {
		valMap[val.key] = val.value
	}

	err := h.c.HSet(context.Background(), key, valMap).Err()
	return err
}

func (h *HashesRepo) GetFieldValue(key string, field string) (string, error) {
	res, err := h.c.HGet(context.Background(), key, field).Result()
	if err != nil {
		return "", err
	}

	return res, nil
}

func (h *HashesRepo) GetAllFieldsValues(ctx context.Context, key string) ([]Pair, error) {
	res, err := h.c.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	pairs := make([]Pair, 0, len(res))
	for k, v := range res {
		pairs = append(pairs, Pair{
			key:   k,
			value: v,
		})
	}

	return pairs, nil
}

func (h *HashesRepo) IsFieldExist(key string, field string) (bool, error) {
	res, err := h.c.HExists(context.Background(), key, field).Result()
	if err != nil {
		return false, err
	}
	return res, nil
}

func (h *HashesRepo) GetFields(key string) ([]string, error) {
	res, err := h.c.HKeys(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *HashesRepo) DeleteField(key string, field ...string) error {
	err := h.c.HDel(context.Background(), key, field...).Err()
	return err
}

func (h *HashesRepo) DeleteAndGetFieldValue(key string, field ...string) ([]string, error) {
	res, err := h.c.HGetDel(context.Background(), key, field...).Result()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *HashesRepo) GetAllValues(key string) ([]string, error) {
	res, err := h.c.HVals(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	return res, nil
}
