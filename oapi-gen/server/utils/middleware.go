package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
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
				log.Printf("client: %s, req: %s\n", r.RemoteAddr, string(data))
			} else {
				log.Println("DumpRequest:", err.Error())
			}
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			rec := httptest.NewRecorder()

			next.ServeHTTP(rec, r)

			res := rec.Result()

			// this copies the recorded response to the response writer
			for k, v := range rec.HeaderMap {
				w.Header()[k] = v
			}
			// write http status code
			w.WriteHeader(rec.Code)
			// write body
			body, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			w.Write(body)

			log.Printf("response: %s", body)
		}

	})
}
