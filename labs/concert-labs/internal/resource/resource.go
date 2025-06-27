package resource

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	DB    *sqlx.DB
	Redis redis.UniversalClient
}
