package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckHealthReturnsNilForHealthyServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	err := checkHealth(server.URL)
	if err != nil {
		t.Fatalf("expected health check to succeed, got error: %v", err)
	}
}

func TestCheckHealthReturnsErrorForUnhealthyServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	err := checkHealth(server.URL)
	if err == nil {
		t.Fatal("expected health check to fail, got nil error")
	}
}

func TestCheckHealthReturnsErrorWhenServerIsUnavailable(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	serverURL := server.URL
	server.Close()

	err := checkHealth(serverURL)
	if err == nil {
		t.Fatal("expected health check to fail when server is unavailable, got nil error")
	}
}