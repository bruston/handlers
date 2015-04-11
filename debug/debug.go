package debug

import (
	"log"
	"net/http"
	"time"
)

type debugResponsewriter struct {
	http.ResponseWriter
	code int
}

func (d debugResponsewriter) WriteHeader(n int) {
	d.ResponseWriter.WriteHeader(n)
	d.code = n
}

func New(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		drw := debugResponsewriter{w, 0}
		start := time.Now()
		h.ServeHTTP(drw, r)
		log.Printf("%s - %s - %s - %s - %d - %s", time.Since(start).String(), r.RemoteAddr, r.Method, r.RequestURI, drw.code, r.UserAgent())
	})
}
