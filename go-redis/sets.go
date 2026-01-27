package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type SetsRepo struct {
	c *redis.Client
}

func NewSetsRepo() (*SetsRepo, error) {
	client, err := NewRedisClient()
	if err != nil {
		return nil, err
	}

	return &SetsRepo{
		c: client,
	}, nil
}

func (s *SetsRepo) AddMembers(key string, members ...string) error {
	err := s.c.SAdd(context.Background(), key, members).Err()
	return err
}

func (s *SetsRepo) GetMembers(key string) ([]string, error) {
	res, err := s.c.SMembers(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SetsRepo) RemoveMembers(key string, members ...string) error {
	err := s.c.SRem(context.Background(), key, members).Err()
	return err
}

func (s *SetsRepo) Union(keys ...string) ([]string, error) {
	res, err := s.c.SUnion(context.Background(), keys...).Result()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SetsRepo) IsMember(key string, member string) (bool, error) {
	res, err := s.c.SIsMember(context.Background(), key, member).Result()
	if err != nil {
		return false, err
	}

	return res, nil
}
