package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type ScoreValue struct {
	member string
	score  float64
}

type SortedSetsRepo struct {
	c *redis.Client
}

func NewSortedSetsRepo() (*SortedSetsRepo, error) {
	client, err := NewRedisClient()
	if err != nil {
		return nil, err
	}

	return &SortedSetsRepo{
		c: client,
	}, nil
}

func (s *SortedSetsRepo) AddMember(key string, value string, score float64) error {
	err := s.c.ZAdd(context.Background(), key, redis.Z{
		Score:  score,
		Member: value,
	}).Err()

	return err
}

func (s *SortedSetsRepo) AddMembers(key string, value []ScoreValue) error {
	pipe := s.c.Pipeline()
	for _, val := range value {
		pipe.ZAdd(context.Background(), key, redis.Z{
			Score:  val.score,
			Member: val.member,
		})
	}

	_, err := pipe.Exec(context.Background())
	return err
}

func (s *SortedSetsRepo) GetRange(key string, start int64, end int64) ([]string, error) {
	res, err := s.c.ZRange(context.Background(), key, start, end).Result()

	s.c.Pipeline()
	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *SortedSetsRepo) GetRangeWithScores(key string, start int64, end int64) ([]redis.Z, error) {
	res, err := s.c.ZRangeWithScores(context.Background(), key, start, end).Result()
	if err != nil {
		return nil, err
	}

	return res, err
}
