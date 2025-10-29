package middleware

import (
	"testing"
	"errors"
	"net/http"
	"net/http/httptest"
	"context"
	"fmt"
)

// TestAuthenticate tests the Authenticate function
func TestAuthenticate(t *testing.T) {
	tests := []struct {
		username string
		password string
		expected string
		err      error
	}{
		{"user1", "password1", "example", nil},
		{"", "password1", "", errors.New("no username")},
		{"user2", "", "example", nil},
	}

	for _, tt := range tests {
		t.Run(tt.username, func(t *testing.T) {
			result, err := Authenticate(tt.username, tt.password)
			if (err != nil) && err.Error() != tt.err.Error() {
				t.Errorf("expected error %v, got %v", tt.err, err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

type mockHandler struct{}

func (m *mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// A simple mock handler that just returns 200 OK
	w.WriteHeader(http.StatusOK)
}

func TestBasicAuth_ValidCredentials(t *testing.T) {
	// Prepare the test
	handler := &mockHandler{}
	middleware := BasicAuth(handler)

	// Create a request with basic auth credentials
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("user1", "password1") // valid credentials

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the middleware
	middleware.ServeHTTP(rr, req)

	// Check if the request passed through to the handler
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", status)
	}
}

func TestBasicAuth_InvalidCredentials(t *testing.T) {
	// Prepare the test
	handler := &mockHandler{}
	middleware := BasicAuth(handler)

	// Create a request with invalid basic auth credentials
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("user1", "wrongpassword") // invalid credentials

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the middleware
	middleware.ServeHTTP(rr, req)

	// Check if the middleware rejected the request with 401 Unauthorized
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("expected status 401 Unauthorized, got %v", status)
	}

	// Check if the "WWW-Authenticate" header is set
	if header := rr.Header().Get("WWW-Authenticate"); header == "" {
		t.Error("expected WWW-Authenticate header to be set")
	}
}

func TestBasicAuth_MissingCredentials(t *testing.T) {
	// Prepare the test
	handler := &mockHandler{}
	middleware := BasicAuth(handler)

	// Create a request without basic auth credentials
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the middleware
	middleware.ServeHTTP(rr, req)

	// Check if the middleware rejected the request with 401 Unauthorized
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("expected status 401 Unauthorized, got %v", status)
	}

	// Check if the "WWW-Authenticate" header is set
	if header := rr.Header().Get("WWW-Authenticate"); header == "" {
		t.Error("expected WWW-Authenticate header to be set")
	}
}

func TestBasicAuth_UserContext(t *testing.T) {
	// Prepare the test
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the user from context
		user, ok := r.Context().Value(userKey).(*User)
		if !ok {
			t.Fatal("expected user to be set in context")
		}
		if user.Name != "user1" {
			t.Errorf("expected username 'user1', got '%s'", user.Name)
		}
		if user.Access != "example" {
			t.Errorf("expected access 'example', got '%s'", user.Access)
		}
		w.WriteHeader(http.StatusOK)
	})
	middleware := BasicAuth(handler)

	// Create a request with basic auth credentials
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("user1", "password1") // valid credentials

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the middleware
	middleware.ServeHTTP(rr, req)

	// Check if the request passed through to the handler
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", status)
	}
}
