package handlers

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: handle arguments, auth

// var favicon, _ = config.Embedded.ReadFile("favicon.ico")
var favicon = []byte{}

var handlerTests = []struct {
	name         string
	method       string
	path         string
	function     http.HandlerFunc
	responseCode int
	expected     string
	base64       bool
}{
	{"favicon", "GET", "/favicon.ico", FaviconHandler, 200, base64.StdEncoding.EncodeToString(favicon), true},
	{"health-check", "GET", "/health-check", HealthCheckHandler, 200, `{"alive": true}`, false},
	{"unmapped", "GET", "/banana", UnmappedHandler, 200, "Unmapped path: /banana", false},
	{"v1-good", "GET", "/", RootHandler1, 200, "Using v1 api", false},
	{"v1-bad", "GET", "/test", RootHandler1, 404, "404 page not found\n", false},
	{"v2-good", "GET", "/", RootHandler2, 200, "Using v2 api", false},
	{"v2-bad", "GET", "/test", RootHandler2, 404, "404 page not found\n", false},
	{"version", "GET", "/version", VersionHandler, 200, config.Version.Full, false},
}

func TestHandlers(t *testing.T) {
	for _, h := range handlerTests {
		t.Run(h.name, func(t *testing.T) {
			req, err := http.NewRequest(h.method, h.path, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.function)
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != h.responseCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, h.responseCode)
			}
			var received string
			if h.base64 {
				received = base64.StdEncoding.EncodeToString(rr.Body.Bytes())
			} else {
				received = rr.Body.String()
			}
			if received != h.expected {
				t.Errorf("handler returned unexpected body: got %v want %v", received, h.expected)
			}
		})
	}
}
