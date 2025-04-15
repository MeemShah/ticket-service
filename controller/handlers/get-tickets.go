package handlers

import (
	"net/http"
	"ticket-service/controller/utils"
)

type GetTicketsReq struct {
	Category  string `json:"category"`
	Price     string `json:"price"`
	Status    string `json:"status"`
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"`
}

func (handlers *Handlers) GetTickets(w http.ResponseWriter, r *http.Request) {
	getTicketsReq := utils.GetPaginationParams(r)

	tickets, err := handlers.ticketSvc.GetTickets(r.Context(), getTicketsReq)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "failed to get tickets", nil)
		return
	}

	utils.SendData(w, tickets)
}
