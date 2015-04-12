package gzipped

import (
	"compress/gzip"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var payload = []byte("hi")

func dummyHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(payload)
	})
}

func TestGzippedSupported(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	req.Header.Set("Accept-Encoding", "gzip")
	h := New(dummyHandler())
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	encoding := w.Header().Get("Content-Encoding")
	if encoding != "gzip" {
		t.Errorf("expecting Content-Encoding header to be: gzip, got %s", encoding)
	}
	vary := w.Header().Get("Vary")
	if vary != "Accept-Encoding" {
		t.Errorf("expecing Vary header to be: Accept-Encoding, got %s", vary)
	}
	gzr, err := gzip.NewReader(w.Body)
	if err != nil {
		t.Fatalf("unable to make gzip reader from response body: %s", err)
	}
	defer gzr.Close()
	body, _ := ioutil.ReadAll(gzr)
	if !reflect.DeepEqual(body, payload) {
		t.Errorf("expecting body to be: %s, got %s", payload, body)
	}
}

func TestGzippedNotSupported(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	w := httptest.NewRecorder()
	h := New(dummyHandler())
	h.ServeHTTP(w, req)
	if !reflect.DeepEqual(w.HeaderMap, make(http.Header)) {
		t.Errorf("expecting Header to be empty, got %v", w.HeaderMap)
	}
	gzr, _ := gzip.NewReader(w.Body)
	if gzr != nil {
		t.Errorf("should not be able to make a gzip Reader from response body")
		defer gzr.Close()
		if _, err := ioutil.ReadAll(gzr); err == nil {
			t.Errorf("should not be able to read a gzipped response")
		}
	}
}
