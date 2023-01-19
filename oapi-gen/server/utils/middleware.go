package utils

import (
	"log"
	"net/http"
	"net/http/httputil"
)

// type MiddlewareFunc func(http.Handler) http.Handler
// must register before  body-reading method
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if (r.URL.Path == "/" && r.Method == "GET") || r.Method == "HEAD" {
			next.ServeHTTP(w, r)
		} else {
			// Do stuff here
			data, err := httputil.DumpRequest(r, true)

			if err == nil {
				log.Println("client:"+r.RemoteAddr, " req:", string(data))
			} else {
				log.Println("DumpRequest:", err.Error())
			}
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		}

	})
}
