package headers

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAddHeaders(t *testing.T) {
	added := http.Header{
		"test": {"1", "2", "3"},
		"1":    {"2", "3"},
	}
	expected := http.Header{
		"1":    {"2", "3"},
		"test": {"1", "2", "3"},
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	h := New(dummyHandler, added)
	h.ServeHTTP(w, req)
	if !reflect.DeepEqual(expected, w.Header()) {
		t.Errorf("expected: %v\n received: %v", expected, w.Header())
	}
}
