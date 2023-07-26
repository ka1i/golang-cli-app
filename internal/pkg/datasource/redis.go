package datasource

import (
	"context"
	"strings"
	"time"

	"github.com/ka1i/cli/internal/pkg/config"
	"github.com/ka1i/cli/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	ctx    context.Context
	client redis.UniversalClient
}

func (r *redisClient) GetClient() (redis.UniversalClient, context.Context) {
	return r.client, r.ctx
}

// Redis
func (r *redisClient) openRedis() {
	options := config.Cfg.Get()

	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    options.Redis.Addrs,
		Password: options.Redis.Passwd,
		DB:       0,

		MasterName: options.Redis.MasterName,

		PoolSize:     10,
		MinIdleConns: 10,

		MaxRetries:      0,
		MinRetryBackoff: 8 * time.Millisecond,
		MaxRetryBackoff: 512 * time.Millisecond,

		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolTimeout:  4 * time.Second,
	})

	r.client = rdb
	r.ctx = context.Background()

	pong, err := r.ping()
	if err != nil {
		panic(err)
	}

	logger.Printf("Redis Use %v [%s]\n", strings.Join(options.Redis.Addrs, ","), pong)

}

func (r *redisClient) ping() (string, error) {
	return r.client.Ping(r.ctx).Result()
}
