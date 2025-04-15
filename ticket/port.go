package ticket

import (
	"context"
	"ticket-service/controller/utils"
	"ticket-service/dto"
	"time"
)

type Service interface {
	CreateTickets(ctx context.Context, req Createticket) error
	GetTicket(ctx context.Context, ticketId string) (*dto.Ticket, error)
	GetTickets(ctx context.Context, filterBy utils.PaginationParams) ([]*dto.Ticket, error)
	CancleView(ctx context.Context, ticketId string) bool
	ConfirmTicket(ctx context.Context, req ConfirmTicket) error
}

type Cache interface {
	SetNX(ctx context.Context, lockKey, lockValue string, lockTTL time.Duration) (bool, error)
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, lockTTL time.Duration) error
	Del(ctx context.Context, key string) bool
}

type TicketRepo interface {
	Create(ctx context.Context, limit int, category string, price int, status string) error
	Get(ctx context.Context, ticketId string) (*dto.Ticket, error)
	Update(ctx context.Context, columns []string, values []any, ticketId string) error
	GetTickets(ctx context.Context, filterBy utils.PaginationParams) ([]*dto.Ticket, error)
}
