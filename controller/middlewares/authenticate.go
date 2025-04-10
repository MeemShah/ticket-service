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
		m.Cache.SetNX(r.Context(), lockKey, lockValue, lockTTL)

		// release lock after function execution if locked
		defer func() {
			val, err := m.Cache.Get(r.Context(), lockKey)
			if err == nil && val == lockValue {
				_ = m.Cache.Del(r.Context(), lockKey)
			}
		}()

		tucketKey := KeyPrifix + ticketID
		status, err := m.Cache.Get(r.Context(), tucketKey)
		if err != nil {
			slog.Error("failed to get ticket", logger.Extra(map[string]any{
				"error":     err.Error(),
				"ticketKey": tucketKey,
			}))
			utils.SendError(w, http.StatusBadRequest, "Something went wrong", nil)
			return
		}

		// reject request if processed by others
		if status != StatusAvailable {
			utils.SendError(w, http.StatusBadRequest, "already checking by others", nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}
