package middlewares

import (
	"log"
	"log/slog"
	"net/http"
	"ticket-service/controller/utils"
	"ticket-service/logger"
	"time"
)

const (
	TicketID   = "ticket-id"
	UserID     = "user-id"
	KeyPrifix  = "ticket: "
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

		// apply lock machanism
		lockKey := LockPrifix + ticketID
		lockTTL := time.Duration(m.cnf.HoldTicktViewInSeconds) * time.Second

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
			log.Println("Someone is already processing this ticket. user-id: ", userID)
			utils.SendError(w, http.StatusTooManyRequests, "Someone is already processing this ticket", nil)
			return
		}

		log.Println("ticket processing. for user-id: ", userID)
		next.ServeHTTP(w, r)
	})
}
