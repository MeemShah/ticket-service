package ticket

import (
	"context"
	"log/slog"
	"ticket-service/logger"
)

type Createticket struct {
	TotalNumberOfTicket int
	Catagory            string
	Price               int
	Status              string
}

func (s *service) CreateTickets(ctx context.Context, req Createticket) error {
	err := s.ticketRepo.Create(ctx, req.TotalNumberOfTicket, req.Catagory, req.Price, req.Status)
	if err != nil {
		slog.Error("failed to create tickets", logger.Extra(map[string]any{
			"error": err.Error(),
			"reqq":  req,
		}))

		return err
	}

	return nil
}
