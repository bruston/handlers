package headers

import "net/http"

// New wraps a http.Handler, adding a collection of headers to the response.
func New(h http.Handler, headers http.Header) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range headers {
			w.Header()[k] = append(w.Header()[k], v...)
		}
		h.ServeHTTP(w, r)
	})
}

// TODO: add some commonly used header maps to pass to New(), Strict-Transport-Security etc
