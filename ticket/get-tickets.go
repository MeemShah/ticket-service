package ticket

import (
	"context"
	"ticket-service/controller/utils"
	"ticket-service/dto"
)

func (s *service) GetTickets(ctx context.Context, filterParams utils.PaginationParams) ([]*dto.Ticket, error) {
	tickets, err := s.ticketRepo.GetTickets(ctx, filterParams)
	if err != nil {
		return nil, err
	}
	
	return tickets, nil
}
