package middlewares

import (
	"log/slog"
	"net/http"
	"ticket-service/controller/utils"
	"ticket-service/logger"
	"time"

	"github.com/google/uuid"
)

const (
	TicketID   = "ticket-id"
	KeyPrifix  = "ticket: "
	LockPrifix = "lock: "

	StatusAvailable  = "AVAILABLE"
	StatusInProgress = "IN_PROGRESS"
	StatusBooked     = "BOOKED"
)

func (m *Middleware) AuthenticateTicket(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ticketID := r.Header.Get(TicketID)
		if ticketID == "" {
			utils.SendError(w, http.StatusUnauthorized, "Unauthorized: Missing ticket id", nil)
			return
		}

		// apply lock machanism
		lockKey := LockPrifix + ticketID
		lockValue := uuid.New().String()
		lockTTL := time.Duration(m.cnf.HoldTicktViewInSeconds) * time.Second

		ok, err := m.Cache.SetNX(r.Context(), lockKey, lockValue, lockTTL)
		if err != nil {
			slog.Error("failed to acquire lock", logger.Extra(map[string]any{
				"error":     err.Error(),
				"ticket_id": ticketID,
				"lock_key":  lockKey,
			}))

			utils.SendError(w, http.StatusInternalServerError, "Internal Server Error: Could not acquire lock", nil)
			return
		}

		if !ok {
			utils.SendError(w, http.StatusTooManyRequests, "Someone is already processing this ticket", nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}
