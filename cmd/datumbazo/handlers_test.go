package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TODO: convert to table of tests since they're all the same
// request path
// response code
// body
// text or base64 of bytes

func TestFaviconHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/favicon.ico", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(faviconHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected, _ := embedded.ReadFile("favicon.ico")

	// TODO: base64 string for favicon

	if !bytes.Equal(rr.Body.Bytes(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUnmappedHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/banana", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(unmappedHandler)
	handler.ServeHTTP(rr, req)
	wantStatus := http.StatusOK
	gotStatus := rr.Code
	if wantStatus != gotStatus {
		t.Errorf("handler returned wrong status code: got %v want %v", gotStatus, wantStatus)
	}
	wantBody := `Unmapped path: /banana`
	gotBody := rr.Body.String()
	if wantBody != gotBody {
		t.Errorf("handler returned unexpected body: got %v want %v", gotBody, wantBody)
	}
}

func TestVersionHander(t *testing.T) {
	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(versionHandler)
	handler.ServeHTTP(rr, req)
	wantStatus := http.StatusOK
	gotStatus := rr.Code
	if wantStatus != gotStatus {
		t.Errorf("handler returned wrong status code: got %v want %v", gotStatus, wantStatus)
	}
	wantBody := "Datumbazo Server"
	gotBody := rr.Body.String()
	if !strings.HasPrefix(gotBody, wantBody) {
		t.Errorf("handler returned unexpected body: got %v want %v", gotBody, wantBody)
	}
}
