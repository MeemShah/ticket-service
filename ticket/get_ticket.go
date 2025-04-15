package ticket

import (
	"context"
	"ticket-service/dto"
)

func (s *service) GetTicket(ctx context.Context, ticketId string) (*dto.Ticket, error) {
	ticket, err := s.ticketRepo.Get(ctx, ticketId)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}
