package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"
	"ticket-service/controller/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	UidKey       contextKey = "uid"
	UserRoleKey  contextKey = "role"
	UserEmailKey contextKey = "email"
)

type AuthClaims struct {
	Uid   string `json:"uid"`
	Role  string `json:"role"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func unauthorizedResponse(w http.ResponseWriter, msg string) {
	log.Println("Unauthorized access:", msg)
	utils.SendError(w, http.StatusUnauthorized, "Unauthorized: "+msg, nil)
}
func (m *Middleware) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		var tokenStr string
		if header != "" {
			tokens := strings.Split(header, " ")
			if len(tokens) != 2 || tokens[0] != "Bearer" {
				unauthorizedResponse(w, "Invalid Authorization format")
				return
			}
			tokenStr = tokens[1]
		} else {
			tokenStr = r.URL.Query().Get("auth")
		}
		if tokenStr == "" {
			unauthorizedResponse(w, "Missing token")
			return
		}
		var claims AuthClaims
		token, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(m.cnf.JwtSecret), nil
		})
		if err != nil {
			unauthorizedResponse(w, "Invalid token: "+err.Error())
			return
		}
		if !token.Valid {
			unauthorizedResponse(w, "Token is not valid")
			return
		}
		if claims.ExpiresAt.Time.Before(time.Now()) {
			unauthorizedResponse(w, "Token has expired")
			return
		}

		ctx := context.WithValue(r.Context(), UidKey, claims.Uid)
		ctx = context.WithValue(ctx, UserRoleKey, claims.Role)
		ctx = context.WithValue(ctx, UserEmailKey, claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
