package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisConfig represents Redis configuration
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

var RedisClient *redis.Client

// InitRedis initializes the Redis connection
func InitRedis(config RedisConfig) error {

	// Create Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password:     config.Password,
		DB:           config.DB,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	RedisClient = rdb
	log.Println("Redis connected successfully")
	return nil
}

// GetRedis returns the Redis client instance
func GetRedis() *redis.Client {
	return RedisClient
}

// CloseRedis closes the Redis connection
func CloseRedis() error {
	if RedisClient != nil {
		if err := RedisClient.Close(); err != nil {
			return fmt.Errorf("failed to close Redis connection: %w", err)
		}
		log.Println("Redis connection closed")
	}
	return nil
}

// SetWithExpiration sets a key-value pair with expiration
func SetWithExpiration(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

// Get retrieves a value by key
func Get(ctx context.Context, key string) (string, error) {
	return RedisClient.Get(ctx, key).Result()
}

// Delete removes a key
func Delete(ctx context.Context, key string) error {
	return RedisClient.Del(ctx, key).Err()
}

// Exists checks if a key exists
func Exists(ctx context.Context, key string) (bool, error) {
	result, err := RedisClient.Exists(ctx, key).Result()
	return result > 0, err
}
