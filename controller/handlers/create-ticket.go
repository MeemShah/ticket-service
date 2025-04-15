package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"ticket-service/controller/utils"
	"ticket-service/logger"
	"ticket-service/ticket"
)

type CreateTicketReq struct {
	TotalNumberOfTicket int    `json:"total_number_of_ticket" validate:"required"`
	Catagory            string `json:"catagory" validate:"required"`
	Price               int    `json:"price" validate:"required"`
	Status              string `json:"status" validate:"required"`
}

func (h *Handlers) Createtickets(w http.ResponseWriter, r *http.Request) {
	var req CreateTicketReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		slog.Error("Failed to decode request body", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		utils.SendError(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = utils.Validate(req)
	if err != nil {
		slog.Error("Invalid request body", logger.Extra(map[string]any{
			"body": req,
		}))
		utils.SendError(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = h.ticketSvc.CreateTickets(r.Context(), ticket.Createticket{
		TotalNumberOfTicket: req.TotalNumberOfTicket,
		Catagory:            req.Catagory,
		Price:               req.Price,
		Status:              req.Status,
	})
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "failed to create tickets", nil)
		return
	}

	utils.SendData(w, "ticket create successfull")

}
