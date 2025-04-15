package ticket

import (
	"context"
	"ticket-service/dto"
)

type Service interface {
	CreateTickets(ctx context.Context, req Createticket) error
	GetTicket(ctx context.Context, ticketId string) (*dto.Ticket, error)
	CancleView(ctx context.Context, ticketId string) bool
}

type Cache interface {
	Del(ctx context.Context, key string) bool
	Get(ctx context.Context, key string) (string, error)
}

type TicketRepo interface {
	Create(ctx context.Context, limit int, category string, price int, status string) error
	Get(ctx context.Context, ticketId string) (*dto.Ticket, error)
}
