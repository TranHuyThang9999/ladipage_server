package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"ladipage_server/core/adapters"
	"ladipage_server/core/adapters/cache"

	"time"

	"github.com/redis/go-redis/v9"
)

type cacheRepository struct {
	db *adapters.Redis
}

func NewRepositoryCache(db *adapters.Redis) cache.CacheOperations {
	return &cacheRepository{
		db: db,
	}
}

func (c *cacheRepository) Delete(ctx context.Context, key string) error {
	err := c.db.Client().Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *cacheRepository) Exists(ctx context.Context, key string) (bool, error) {
	result, err := c.db.Client().Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

func (c *cacheRepository) Expire(ctx context.Context, key string, expiration time.Duration) error {
	err := c.db.Client().Expire(ctx, key, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *cacheRepository) Get(ctx context.Context, key string, dest interface{}) error {
	val, err := c.db.Client().Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}
	err = json.Unmarshal([]byte(val), dest)
	if err != nil {
		return err
	}
	return nil
}

func (c *cacheRepository) HGet(ctx context.Context, key string, field string) (string, error) {
	val, err := c.db.Client().HGet(ctx, key, field).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}
	return val, nil
}

func (c *cacheRepository) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	result, err := c.db.Client().HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *cacheRepository) HSet(ctx context.Context, key string, values ...interface{}) error {
	hashValues := make(map[string]interface{})
	for i := 0; i < len(values); i += 2 {
		hashValues[values[i].(string)] = values[i+1].(string)
	}

	err := c.db.Client().HSet(ctx, key, hashValues).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *cacheRepository) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value to JSON: %w", err)
	}

	err = c.db.Client().Set(ctx, key, jsonBytes, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set value in Redis: %w", err)
	}

	return nil
}
