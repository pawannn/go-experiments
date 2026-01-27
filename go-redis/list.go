package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type ListRepo struct {
	c *redis.Client
}

func NewListRepo() (*ListRepo, error) {
	client, err := NewRedisClient()
	if err != nil {
		return nil, err
	}

	return &ListRepo{
		c: client,
	}, nil
}

func (l *ListRepo) AddElementToLeft(key string, elements ...interface{}) error {
	err := l.c.LPush(context.Background(), key, elements...).Err()
	return err
}

func (l *ListRepo) AddElementToRight(key string, elements ...interface{}) error {
	err := l.c.RPush(context.Background(), key, elements...).Err()
	return err
}

func (l *ListRepo) GetRangeElements(key string, start int64, end int64) ([]string, error) {
	res, err := l.c.LRange(context.Background(), key, start, end).Result()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (l *ListRepo) PopLeft(key string) error {
	err := l.c.LPop(context.Background(), key).Err()
	return err
}

func (l *ListRepo) PopRight(key string) error {
	err := l.c.RPop(context.Background(), key).Err()
	return err
}

func (l *ListRepo) Length(key string) (int64, error) {
	res, err := l.c.LLen(context.Background(), key).Result()
	if err != nil {
		return -1, err
	}

	return res, nil
}

func (l *ListRepo) ElementAtIndex(key string, idx int64) (string, error) {
	res, err := l.c.LIndex(context.Background(), key, idx).Result()
	if err != nil {
		return "", err
	}

	return res, nil
}

func (l *ListRepo) IndexOfElement(key string, element string) (int64, error) {
	res, err := l.c.LPos(context.Background(), key, element, redis.LPosArgs{
		Rank:   1,  // occourance position 1, 2, 3, 4, -2, -1
		MaxLen: 10, // Limit how many list elements Redis scans from the start
	}).Result()

	if err != nil {
		return -1, err
	}

	return res, nil
}

func (l *ListRepo) Insert(key string, op string, pivot string, newElement string) error {
	err := l.c.LInsert(context.Background(), key, op, pivot, newElement).Err()
	return err
}

func (l *ListRepo) InsertAfter(key string, pivot string, newElement string) error {
	err := l.c.LInsertAfter(context.Background(), key, pivot, newElement).Err()
	return err
}

func (l *ListRepo) InsertBefore(key string, pivot string, newElement string) error {
	err := l.c.LInsertBefore(context.Background(), key, pivot, newElement).Err()
	return err
}
