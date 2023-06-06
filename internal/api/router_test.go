package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type mockNavigator struct{}

func (m *mockNavigator) GetRoute(source, destination *Location) (Route, error) {
	// Mock implementation for GetRoute
	return Route{}, nil
}

func TestHandleRoutesRequest(t *testing.T) {
	// Create a new API instance with a mock navigator
	api := Initialize(&mockNavigator{})

	// Create a test request
	req, err := http.NewRequest("GET", "/routes?src=10,20&dst=30,40&dst=50,60", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test response recorder
	recorder := httptest.NewRecorder()

	// Call the handler function
	api.handleRoutesRequest(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", recorder.Code, http.StatusOK)
	}

}

func TestParseRequestURL(t *testing.T) {
	// Test valid request URL
	requestValid := &url.URL{
		RawQuery: "src=10,20&dst=30,40&dst=50,60",
	}
	source, destinations, err := parseRequestURL(requestValid)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Add assertions for the parsed values
	if source.Lat != 10 {
		t.Errorf("unexpected source latitude: got %v, want %v", source.Lat, 10)
	}

	if len(destinations) != 2 {
		t.Errorf("unexpected number of destinations: got %v, want %v", len(destinations), 2)
	}

	// Test invalid request URL
	requestInvalid := &url.URL{
		RawQuery: "src=10,20&dst=invalid",
	}
	_, _, err = parseRequestURL(requestInvalid)
	if err == nil {
		t.Error("expected an error, but got nil")
	}

}
