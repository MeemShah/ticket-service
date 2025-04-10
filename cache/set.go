package cache

import (
	"context"
	"time"
)

func (c *cache) Set(ctx context.Context, key string, value string, lockTTL time.Duration) error {
	err := c.writeClient.Set(ctx, key, value, lockTTL).Err()
	if err != nil {
		return err
	}

	return nil
}
