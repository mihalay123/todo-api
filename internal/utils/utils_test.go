package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseIDFromPath(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/todos/42", nil)
	id, err := ParseIDFromPath(req, "/todos/")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != 42 {
		t.Errorf("expected 42, got %d", id)
	}
}

func TestParseIDFromPath_Invalid(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/todos/abc", nil)
	_, err := ParseIDFromPath(req, "/todos/")
	if err == nil {
		t.Error("expected error for non-integer ID, got nil")
	}
}
