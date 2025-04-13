package handlers

import (
	"net/http"
	"ticket-service/controller/utils"
)

const TicketID = "ticket-id"

func (h *Handlers) CancleView(w http.ResponseWriter, r *http.Request) {
	ticketId := r.Header.Get(TicketID)

	if ticketId == "" {
		utils.SendError(w, http.StatusBadRequest, "ticket id required", nil)
		return
	}

	ok := h.ticketSvc.CancleView(r.Context(), ticketId)
	if !ok {
		utils.SendError(w, http.StatusInternalServerError, "failed to parform cancle view", nil)
		return
	}

	utils.SendData(w, "success")

}
