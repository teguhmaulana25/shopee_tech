package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Logger code explain itself
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Debugf(
			"%s\t%s\t%s\t",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
