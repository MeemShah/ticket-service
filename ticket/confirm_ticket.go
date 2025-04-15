package ticket

import (
	"context"
	"log/slog"
	"sync"
	"ticket-service/logger"
)

const (
	Column_PurchasedBy = "purchased_by"
	Column_Status      = "status"
)

type ConfirmTicket struct {
	TickitID    string
	PurchasedBy string
	Status      string
}

func (s *service) ConfirmTicket(ctx context.Context, req ConfirmTicket) error {
	var wg sync.WaitGroup
	errCh := make(chan error, 2)

	columns := []string{Column_PurchasedBy, Column_Status}
	values := []any{req.PurchasedBy, req.Status}

	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := s.ticketRepo.Update(ctx, columns, values, req.TickitID); err != nil {
			slog.Error("failed to update database", logger.Extra(map[string]any{
				"error": err.Error(),
				"req":   req,
			}))
			errCh <- err
		}
	}()

	go func() {
		defer wg.Done()
		cacheKey := LockPrifix + req.TickitID
		if err := s.cache.Set(ctx, cacheKey, req.PurchasedBy, 0); err != nil {
			slog.Error("failed to update cache", logger.Extra(map[string]any{
				"error": err.Error(),
				"req":   req,
			}))
			errCh <- err
		}
	}()

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}
