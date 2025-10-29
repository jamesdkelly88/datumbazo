package middleware

import (
	"testing"
)

func TestAuthenticate(t *testing.T) {
	acc, err := Authenticate("username", "password")
	if acc == "" {
		t.Error("Access is empty, expected 'example'")
	}
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
func TestAuthenticateFails(t *testing.T) {
	acc, err := Authenticate("", "")
	if acc != "" {
		t.Errorf("Access is %s, expected empty", acc)
	}
	if err == nil {
		t.Error("Expected an error, got nothing")
	}
}
