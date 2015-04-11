package xhttpmethod

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type dummyHandler struct{}

func (d dummyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func TestMethodOverride(t *testing.T) {
	testTable := []struct {
		reqMethod string
		override  string
		expected  string
	}{
		{"POST", "PATCH", "PATCH"},
		{"POST", "PUT", "PUT"},
		{"POST", "DELETE", "DELETE"},
		{"POST", "HEAD", "POST"},
		{"POST", "GET", "POST"},
		{"POST", "", "POST"},
		{"HEAD", "PUT", "HEAD"},
		{"GET", "PUT", "GET"},
		{"DELETE", "PUT", "DELETE"},
		{"PUT", "DELETE", "PUT"},
		{"OPTIONS", "DELETE", "OPTIONS"},
	}

	h := New(dummyHandler{})
	for i, tt := range testTable {
		w := httptest.NewRecorder()
		r, err := http.NewRequest(tt.reqMethod, "/", nil)
		if err != nil {
			t.Errorf("%d. unexpected error: %s", i, err)
		}
		r.Header.Set("X-HTTP-Method-Override", tt.override)
		h.ServeHTTP(w, r)
		if r.Method != tt.expected {
			t.Errorf("%d. expecting method %s, got %s", i, tt.expected, r.Method)
		}
	}
}
