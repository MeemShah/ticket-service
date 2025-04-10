package middlewares

import (
	"context"
	"ticket-service/config"
	"time"
)

type Cache interface {
	SetNX(ctx context.Context, lockKey, lockValue string, lockTTL time.Duration) (bool, error)
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, lockTTL time.Duration) error
	Del(ctx context.Context, key string) bool
}

type Middleware struct {
	cnf   *config.Config
	Cache Cache
}

func NewMiddleware(cnf *config.Config, cache Cache) *Middleware {
	return &Middleware{
		cnf:   cnf,
		Cache: cache,
	}
}
