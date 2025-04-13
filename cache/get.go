package cache

import (
	"context"
)

const (
	StatusAvailable  = "AVAILABLE"
	StatusInProgress = "IN_PROGRESS"
	StatusBooked     = "BOOKED"
)

func (c *cache) Get(ctx context.Context, key string) (string, error) {
	data, err := c.readClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return data, nil
}
