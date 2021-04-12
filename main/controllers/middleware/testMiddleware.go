package middleware

import (
	"log"
	"net/http"
)

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,POST,DELETE,PUT,OPTIONS")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

// w.Header().Set("Content-Type", "application/json")
// w.Header().Set("Access-Control-Allow-Origin", "*")
// w.Header().Set("Access-Control-Allow-Headers", "*")
// w.Header().Set("Access-Control-Expose-Headers", "*")
// w.Header().Set("Access-Control-Expose-Headers", "Authorization")
// w.Header().Set("Access-Control-Allow-Methods", "*")
// w.Header().Set("Access-Control-Allow-Method", "POST")
