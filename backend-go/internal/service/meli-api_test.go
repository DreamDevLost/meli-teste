package service_test

import (
	"backend-api/internal/service"
	"backend-api/internal/service/responses"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMeliAPI_GetItemByID(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL and respond accordingly
		if r.URL.Path == "/items/123" {
			// Return a successful response with a JSON payload
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id": "123", "title": "Test Item"}`))
		} else if r.URL.Path == "/items/456" {
			// Return a not found response
			w.WriteHeader(http.StatusNotFound)
		} else {
			// Return an unexpected response
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
	defer server.Close()

	// Create a new MeliAPI instance with the mock HTTP client
	api := service.NewMeliAPI(server.Client())

	// Test case 1: Successful response
	response, err := api.GetItemByID(context.Background(), "123")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expectedResponse := &responses.GetItemResponse{
		ID:    "123",
		Title: "Test Item",
	}
	if response.ID != expectedResponse.ID || response.Title != expectedResponse.Title {
		t.Errorf("unexpected response: got %v, want %v", response, expectedResponse)
	}

	// Test case 2: Not found response
	_, err = api.GetItemByID(context.Background(), "456")
	if !errors.Is(err, service.ErrNotFound) {
		t.Errorf("unexpected error: got %v, want %v", err, service.ErrNotFound)
	}

	// Test case 3: Unexpected response
	_, err = api.GetItemByID(context.Background(), "789")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
