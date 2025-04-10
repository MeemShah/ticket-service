package cache

import (
	"context"
	"fmt"
	"time"
)

func (c *cache) SetNX(ctx context.Context, lockKey, lockValue string, lockTTL time.Duration) (bool, error) {
	ok, err := c.readClient.SetNX(ctx, lockKey, lockValue, lockTTL).Result()
	if err != nil {
		return false, err
	}
	if !ok {
		return false, fmt.Errorf("key is locked, try again later")
	}
	
	return true, nil
}
