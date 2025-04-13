package handlers

import (
	"net/http"
	"ticket-service/controller/utils"
)

func (handlers *Handlers) GetTicket(w http.ResponseWriter, r *http.Request) {
	ticketID := r.Header.Get(TicketID)

	ticket, err := handlers.ticketSvc.GetTicket(r.Context(), ticketID)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "failed to get ticket details", nil)
		return
	}

	utils.SendData(w, ticket)
}
