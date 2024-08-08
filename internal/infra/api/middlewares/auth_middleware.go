package middlewares

import (
	"net/http"
	"strings"

	"github.com/julianojj/fasttrip/internal/infra/adapters"
)

type AuthMiddleware struct {
	sign adapters.Sign
}

func NewAuthMiddleware(
	sign adapters.Sign,
) *AuthMiddleware {
	return &AuthMiddleware{
		sign,
	}
}

func (a *AuthMiddleware) ApplyHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}
		bearerToken := strings.Split(authorization, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		token := bearerToken[1]
		if a.sign.Verify(token) {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
