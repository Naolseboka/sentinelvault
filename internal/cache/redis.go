package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Redis wraps a go-redis client with JSON helpers.
type Redis struct {
	client *redis.Client
}

// OpenRedis connects using a redis URL (redis://host:port/db).
func OpenRedis(url string) (*Redis, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("parse redis url: %w", err)
	}
	client := redis.NewClient(opts)
	return &Redis{client: client}, nil
}

// Client returns the underlying redis client.
func (r *Redis) Client() *redis.Client { return r.client }

// Ping verifies connectivity.
func (r *Redis) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

// Close closes the client.
func (r *Redis) Close() error { return r.client.Close() }

// Get retrieves a string value.
func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// Set stores a string value with TTL.
func (r *Redis) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

// Del removes keys.
func (r *Redis) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

// GetJSON unmarshals JSON into dest.
func (r *Redis) GetJSON(ctx context.Context, key string, dest any) error {
	raw, err := r.Get(ctx, key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(raw), dest)
}

// SetJSON marshals v and stores with TTL.
func (r *Redis) SetJSON(ctx context.Context, key string, v any, ttl time.Duration) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return r.Set(ctx, key, string(data), ttl)
}
