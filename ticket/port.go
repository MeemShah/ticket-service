package ticket

import (
	"context"
	"ticket-service/dto"
)

type Service interface {
	GetTicket(ctx context.Context, ticketId string) (*dto.Ticket, error)
	CancleView(ctx context.Context, ticketId string) bool
}

type Cache interface {
	Del(ctx context.Context, key string) bool
	Get(ctx context.Context, key string) (string, error)
}
