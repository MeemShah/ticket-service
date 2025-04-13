package cache

import (
	"ticket-service/controller/middlewares"
	"ticket-service/ticket"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	middlewares.Cache
	ticket.Cache
}

type cache struct {
	readClient  *redis.Client
	writeClient *redis.Client
}

func NewCache(readClient, writeClient *redis.Client) Cache {
	return &cache{
		readClient:  readClient,
		writeClient: writeClient,
	}
}
