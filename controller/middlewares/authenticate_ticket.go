package middlewares

import (
	"context"
	"log/slog"
	"net/http"
	"ticket-service/controller/utils"
	"ticket-service/logger"
	"time"
)

const (
	TicketIDKey contextKey = "ticket-id"
	UserIdKey   contextKey = "user-id"
)

const (
	TicketID   = "ticket-id"
	UserID     = "user-id"
	LockPrifix = "lock: "

	StatusAvailable  = "AVAILABLE"
	StatusInProgress = "IN_PROGRESS"
	StatusBooked     = "BOOKED"
)

func (m *Middleware) AuthenticateTicket(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ticketID := r.Header.Get(TicketID)
		userID := r.Header.Get(UserID)
		if ticketID == "" || userID == "" {
			utils.SendError(w, http.StatusUnauthorized, "Unauthorized: Missing ticket id", nil)
			return
		}

		lockKey := LockPrifix + ticketID
		lockTTL := time.Duration(m.cnf.HoldTicktViewInSeconds) * time.Second

		bookedBy, err := m.Cache.Get(r.Context(), lockKey)
		if err == nil {
			if bookedBy == userID {
				ctx := context.WithValue(r.Context(), TicketIDKey, ticketID)
				ctx = context.WithValue(ctx, UserIdKey, userID)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}

		ok, err := m.Cache.SetNX(r.Context(), lockKey, userID, lockTTL)
		if err != nil {
			slog.Error("failed to acquire lock", logger.Extra(map[string]any{
				"error":     err.Error(),
				"ticket_id": ticketID,
				"lock_key":  lockKey,
				"host":      m.cnf.ReadRedisUrl,
			}))

			utils.SendError(w, http.StatusInternalServerError, "Internal Server Error: Could not acquire lock", nil)
			return
		}

		if !ok {
			utils.SendError(w, http.StatusLocked, "Someone is already viewing this ticket", nil)
			return
		}

		ctx := context.WithValue(r.Context(), TicketIDKey, ticketID)
		ctx = context.WithValue(ctx, UserIdKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
