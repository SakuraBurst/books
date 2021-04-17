package middleware

import (
	"fmt"
	"net/http"

	"github.com/SakuraBurst/books.git/main/models"
	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(rw, "unauthorized", http.StatusUnauthorized)
			return
		} else {
			tokens, err := jwt.Parse(token, models.KeyFunc)
			if err != nil {
				http.Error(rw, "unauthorized", http.StatusUnauthorized)
				return
			}
			if tokens.Valid {
				fmt.Println("protected route")
				next.ServeHTTP(rw, r)
			} else {
				http.Error(rw, "token is expired", http.StatusUnauthorized)
				return
			}
		}
	})
}
