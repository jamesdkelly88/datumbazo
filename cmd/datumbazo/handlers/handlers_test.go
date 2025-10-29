package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jamesdkelly88/datumbazo/internal/config"
)

// TODO: handle arguments, auth

var favicon, _ = Embedded.ReadFile("favicon.ico")

var version = config.GetVersion(true)
var versionJSON, _ = json.Marshal(version)

var handlerTests = []struct {
	name         string
	method       string
	path         string
	function     http.HandlerFunc
	responseCode int
	expected     string
	base64       bool
}{
	{"favicon", "GET", "/favicon.ico", Favicon, 200, base64.StdEncoding.EncodeToString(favicon), true},
	{"healthz", "GET", "/healthz", Health, 204, "", false},
	{"unmapped", "GET", "/banana", Root, 404, "404 page not found\n", false},
	{"root", "GET", "/", Root, 200, "Welcome to Datumbazo\n", false},
	{"version", "GET", "/version", Version(version), 200, string(versionJSON), false},
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
