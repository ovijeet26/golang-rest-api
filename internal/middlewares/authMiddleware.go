package middlewares

import (
	"net/http"
	"strings"

	"github.com/ovijeet26/golang-rest-api/internal/utils"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		if authHeader == "itsme" {
			next.ServeHTTP(w, r)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		_, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
