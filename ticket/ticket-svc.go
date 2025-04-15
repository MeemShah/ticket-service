package ticket

import (
	"ticket-service/config"
)

type service struct {
	cnf        *config.Config
	cache      Cache
	ticketRepo TicketRepo
}

func NewService(cnf *config.Config, cache Cache, ticketRepo TicketRepo) Service {
	return &service{
		cnf:        cnf,
		cache:      cache,
		ticketRepo: ticketRepo,
	}
}
