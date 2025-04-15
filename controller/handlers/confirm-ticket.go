package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"ticket-service/controller/middlewares"
	"ticket-service/controller/utils"
	"ticket-service/logger"
	"ticket-service/ticket"
)

type ConfirmTicketReq struct {
	Status string `json:"status" validate:"required"`
}

func (handlers *Handlers) ConfirmTicket(w http.ResponseWriter, r *http.Request) {
	ticketID := r.Header.Get(string(middlewares.TicketIDKey))
	userId := r.Header.Get(string(middlewares.UserIdKey))

	var req ConfirmTicketReq
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

	err = handlers.ticketSvc.ConfirmTicket(r.Context(), ticket.ConfirmTicket{
		TickitID:    ticketID,
		PurchasedBy: userId,
		Status:      req.Status,
	})
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "failed to get ticket details", nil)
		return
	}

	utils.SendData(w, "ticket Confirmed")
}
