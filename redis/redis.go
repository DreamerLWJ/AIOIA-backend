package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/DreamerLWJ/AIOIA-backend/config"
	"github.com/go-redis/redis/v9"
)

const (
	_defaultMaxRetries = 1
)

type RdsClient struct {
	*redis.Client
}

func NewRedisClientToml(path string) (*RdsClient, error) {
	var rdsConf config.RedisConfig
	_, err := toml.DecodeFile(path, &rdsConf)
	if err != nil {
		return nil, err
	}

	client := NewRedisClient(&rdsConf)
	return client, nil
}

func NewRedisClient(c *config.RedisConfig) *RdsClient {
	opt := redis.Options{
		Addr: c.Addr,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			fmt.Println("Connect Success")
			return nil
		},
		Username:              c.Username,
		Password:              c.Password,
		DB:                    c.DB,
		MaxRetries:            _defaultMaxRetries,
		DialTimeout:           time.Millisecond * time.Duration(c.ConnTimeout),
		ReadTimeout:           time.Millisecond * time.Duration(c.ReadTimeout),
		WriteTimeout:          time.Millisecond * time.Duration(c.WriteTimeout),
		ContextTimeoutEnabled: true,
		PoolSize:              c.MaxActive,
		PoolTimeout:           time.Millisecond * time.Duration(c.PoolTimeout),
		MinIdleConns:          c.MinIdle,
		MaxIdleConns:          c.MaxIdle,
		ConnMaxIdleTime:       time.Millisecond * time.Duration(c.MaxIdleTime),
	}
	innerClient := redis.NewClient(&opt)
	rdsClient := RdsClient{innerClient}
	return &rdsClient
}
