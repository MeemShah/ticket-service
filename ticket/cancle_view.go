package ticket

import (
	"context"
)

const (
	TicketID   = "ticket-id"
	UserID     = "user-id"
	KeyPrifix  = "ticket: "
	LockPrifix = "lock: "

	StatusAvailable  = "AVAILABLE"
	StatusInProgress = "IN_PROGRESS"
	StatusBooked     = "BOOKED"
)

func (s *service) CancleView(ctx context.Context, ticketId string) bool {
	if ticketId == "" {
		return false
	}

	ok, err := s.cache.Get(ctx, LockPrifix+ticketId)
	if err == nil && ok != "" {
		done := s.cache.Del(ctx, LockPrifix+ticketId)
		return done
	}

	return false
}
