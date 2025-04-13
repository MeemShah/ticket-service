package ticket

import (
	"context"
	"ticket-service/dto"
)

func (s *service) GetTicket(ctx context.Context, ticketId string) (*dto.Ticket,error) {

	return &dto.Ticket{
		Id:       ticketId,
		Catagory: "First-class",
		Price:    15000,
		Status:   "Inprogress",
		IsActive: true,
	}, nil

}
