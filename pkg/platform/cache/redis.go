package cache

import (
	"context"
	"time"

	"github.com/cp-Coder/khelo/pkg/utils"
	redis "github.com/redis/go-redis/v9"
)

// RedisClient pointer to the redis client
type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

var redisClient *RedisClient

// InitRedis starts the redis connection
func InitRedis() error {

	redisConnURL, err := utils.ConnectionURLBuilder("redis")
	if err != nil {
		return err
	}
	ctx := context.Background()

	// Set Redis options.
	options := &redis.Options{
		Addr: redisConnURL,
		DB:   0, // use default DB
		// Password: os.Getenv("REDIS_PASSWORD"),
		// DialTimeout:        10 * time.Second,
		// ReadTimeout:        30 * time.Second,
		// WriteTimeout:       30 * time.Second,
		// PoolSize:           10,
		// PoolTimeout:        30 * time.Second,
		// IdleTimeout:        500 * time.Millisecond,
		// IdleCheckFrequency: 500 * time.Millisecond,
		// TLSConfig: &tls.Config{
		// 	InsecureSkipVerify: true,
		// },
	}

	client := redis.NewClient(options)
	redisClient = &RedisClient{
		client: client,
		ctx:    ctx,
	}
	return nil
}

// CloseRedis closes the redis connection
func CloseRedis() {
	redisClient.client.Close()
}

// GetRedisClient returns the redis client
func GetRedisClient() *RedisClient {
	return redisClient
}

// Get gets the value in redis
func (r *RedisClient) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

// Set sets the value in redis
func (r *RedisClient) Set(key string, value interface{}) error {
	return r.client.Set(r.ctx, key, value, 0).Err()
}

// Del deletes the value in redis
func (r *RedisClient) Del(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

// Getx gets the value in redis with expiry
func (r *RedisClient) Getx(key string, expiry time.Duration) (string, error) {
	return r.client.GetEx(r.ctx, key, expiry).Result()
}

// Setx sets the value in redis with expiry
func (r *RedisClient) Setx(key string, value interface{}, expiry time.Duration) error {
	return r.client.Set(r.ctx, key, value, expiry).Err()
}
