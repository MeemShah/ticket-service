package cache

import "context"

func (c *cache) Del(ctx context.Context, key string) bool {
	err := c.writeClient.Del(ctx, key).Err()
	return err == nil
}
