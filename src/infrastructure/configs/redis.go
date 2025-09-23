package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type Redis struct {
	Client *redis.Client
}

func NewRedis(config RedisConfig) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password:     config.Password,
		DB:           config.DB,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MinIdleConns: 5,
	})

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	log.Println("✅ Redis connected successfully")

	return &Redis{Client: rdb}, nil
}

func (r *Redis) Close() error {
	return r.Client.Close()
}

func (r *Redis) GetClient() *redis.Client {
	return r.Client
}

func (r *Redis) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return r.Client.Ping(ctx).Err()
}