package debug

import (
	"log"
	"net/http"
	"time"
)

type debugResponseWriter struct {
	http.ResponseWriter
	code    int
	written int64
}

func (d *debugResponseWriter) WriteHeader(n int) {
	d.ResponseWriter.WriteHeader(n)
	d.code = n
}

func (d *debugResponseWriter) Write(b []byte) (int, error) {
	n, err := d.ResponseWriter.Write(b)
	d.written += int64(n)
	return n, err
}

func New(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		drw := &debugResponseWriter{w, 0, 0}
		start := time.Now()
		h.ServeHTTP(drw, r)
		log.Printf("%s - %s - %s - %s - %d - %d - %s", time.Since(start).String(), r.RemoteAddr, r.Method, r.RequestURI, drw.code, drw.written, r.UserAgent())
	})
}
