package handlers

import (
	"ticket-service/config"
	"ticket-service/ticket"
)

type Handlers struct {
	cnf       *config.Config
	ticketSvc ticket.Service
}

func NewHandlers(cnf *config.Config, ticketSvc ticket.Service) *Handlers {
	return &Handlers{
		cnf:       cnf,
		ticketSvc: ticketSvc,
	}
}
