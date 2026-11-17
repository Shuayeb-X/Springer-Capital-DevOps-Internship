package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test root endpoint
func TestRootHandler(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	rootHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	var response map[string]string

	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Invalid JSON response: %v", err)
	}

	if response["application"] != "Springer Capital Demo App" {
		t.Errorf("Unexpected application name: %s", response["application"])
	}

	if response["status"] != "running" {
		t.Errorf("Expected running, got %s", response["status"])
	}
}

// Test health endpoint
func TestHealthHandler(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	var response map[string]string

	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Invalid JSON response: %v", err)
	}

	if response["status"] != "healthy" {
		t.Errorf("Expected healthy, got %s", response["status"])
	}
}

// Test health endpoint with POST method
func TestHealthHandlerMethodNotAllowed(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, got %d", rec.Code)
	}
}

// Test unknown route
func TestUnknownRoute(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	rec := httptest.NewRecorder()

	newHandler().ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", rec.Code)
	}
}
