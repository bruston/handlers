package xhttpmethod

import "net/http"

func New(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			h.ServeHTTP(w, r)
			return
		}
		m := r.Header.Get("X-HTTP-Method-Override")
		if m == "PATCH" || m == "PUT" || m == "DELETE" {
			r.Method = m
		}
		h.ServeHTTP(w, r)
	})
}
