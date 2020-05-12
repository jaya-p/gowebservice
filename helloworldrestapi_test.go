// run test: go test -v
package gowebservice

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetHandler(t *testing.T) {

	r := httptest.NewRequest("GET", "/api/v1/helloworld", nil)
	w := httptest.NewRecorder()
	h := http.HandlerFunc(HelloworldRestAPIHandler)

	h.ServeHTTP(w, r)

	// check HTTP response status code
	wantCode := http.StatusOK
	if s := w.Code; s != wantCode {
		t.Errorf("handler return wrong status code: got %v, want %v", s, wantCode)
	}

	// check HTTP response header content type
	wantContentType := "text/plain; charset=utf-8"
	if c := w.Header().Get("Content-type"); c != wantContentType {
		t.Errorf("handler return wrong status code: got %v, want %v", c, wantContentType)
	}

	// check HTTP response body
	wantBody := "Hello World"
	if w.Body.String() != wantBody {
		t.Errorf("handler return wrong status code: got %v, want %v", w.Body.String(), wantBody)
	}
}

func TestPostHandler(t *testing.T) {

	r := httptest.NewRequest("POST", "/api/v1/helloworld", strings.NewReader("{\"Name\": \"There\"}"))
	w := httptest.NewRecorder()
	h := http.HandlerFunc(HelloworldRestAPIHandler)

	h.ServeHTTP(w, r)

	// check HTTP response status code
	wantCode := http.StatusOK
	if s := w.Code; s != wantCode {
		t.Errorf("handler return wrong status code: got %v, want %v", s, wantCode)
	}

	// check HTTP response header content type
	wantContentType := "text/plain; charset=utf-8"
	if c := w.Header().Get("Content-type"); c != wantContentType {
		t.Errorf("handler return wrong status code: got %v, want %v", c, wantContentType)
	}

	// check HTTP response body
	wantBody := "Hello World, There!"
	if w.Body.String() != wantBody {
		t.Errorf("handler return wrong status code: got %v, want %v", w.Body.String(), wantBody)
	}
}

func TestDefaultHandler(t *testing.T) {

	r := httptest.NewRequest("DELETE", "/api/v1/helloworld", nil)
	w := httptest.NewRecorder()
	h := http.HandlerFunc(HelloworldRestAPIHandler)

	h.ServeHTTP(w, r)

	// check HTTP response status code
	wantCode := http.StatusNotFound
	if s := w.Code; s != wantCode {
		t.Errorf("handler return wrong status code: got %v, want %v", s, wantCode)
	}

	// check HTTP response header content type
	wantContentType := "text/plain; charset=utf-8"
	if c := w.Header().Get("Content-type"); c != wantContentType {
		t.Errorf("handler return wrong status code: got %v, want %v", c, wantContentType)
	}

	// check HTTP response body
	wantBody := "Your requested method (DELETE) is not found"
	if w.Body.String() != wantBody {
		t.Errorf("handler return wrong status code: got %v, want %v", w.Body.String(), wantBody)
	}
}
