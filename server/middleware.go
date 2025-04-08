package server

import "net/http"

func clientHintsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Accept-CH", "Sec-CH-Prefers-Color-Scheme")
		w.Header().Add("Critical-CH", "Sec-CH-Prefers-Color-Scheme")
		w.Header().Add("Vary", "Sec-CH-Prefers-Color-Scheme")
		next.ServeHTTP(w, r)
	})
}
