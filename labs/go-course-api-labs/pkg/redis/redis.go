package redis

import (
	"time"

	"example.com/pkg/config"
	"github.com/redis/go-redis/v9"
)

func New(c config.Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.GetAddr(),
		Password:     c.Password,
		DB:           c.DB,
		ReadTimeout:  time.Duration(c.ReadTimeoutSec) * time.Second,
		WriteTimeout: time.Duration(c.WriteTimeoutSec) * time.Second,
		DialTimeout:  time.Duration(c.ConnDialTimeoutSec) * time.Second,
		PoolTimeout:  time.Duration(c.ConnPoolTimeoutSec) * time.Second,
		PoolSize:     c.ConnPoolSize,
		MinIdleConns: c.MinIdleConn,
		MaxIdleConns: c.MaxIdleConn,
	})
	return rdb
}
