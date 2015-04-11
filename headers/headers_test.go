package headers

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAddHeaders(t *testing.T) {
	existing := http.Header{
		"existing": {"1"},
		"1":        {"1"},
	}
	added := http.Header{
		"test": {"1", "2", "3"},
		"1":    {"2", "3"},
	}
	expected := http.Header{
		"existing": {"1"},
		"1":        {"1", "2", "3"},
		"test":     {"1", "2", "3"},
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	req.Header = existing

	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	h := New(dummyHandler, added)
	h.ServeHTTP(w, req)
	if !reflect.DeepEqual(expected, req.Header) {
		t.Errorf("expected: %v\n received: %v", expected, req.Header)
	}
}
