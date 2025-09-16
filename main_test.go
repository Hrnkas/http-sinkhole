package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSinkHoleHandler(t *testing.T) {
	tests := []struct {
		name   string
		method string
		path   string
	}{
		{"GET root", "GET", "/"},
		{"POST root", "POST", "/"},
		{"PUT root", "PUT", "/"},
		{"DELETE root", "DELETE", "/"},
		{"GET arbitrary path", "GET", "/some/random/path"},
		{"POST with query params", "POST", "/api/data?param=value"},
		{"PATCH nested path", "PATCH", "/nested/very/deep/path"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()

			sinkHoleHandler(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			// Check status code is always 200
			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
			}

			// Check response body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			expectedBody := "OK\n"
			if string(body) != expectedBody {
				t.Errorf("Expected body %q, got %q", expectedBody, string(body))
			}
		})
	}
}

func TestServerIntegration(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(sinkHoleHandler))
	defer server.Close()

	// Test various HTTP methods and paths
	testCases := []struct {
		method string
		path   string
	}{
		{"GET", "/"},
		{"POST", "/api/test"},
		{"PUT", "/data"},
		{"DELETE", "/item/123"},
		{"HEAD", "/health"},
		{"OPTIONS", "/cors"},
	}

	for _, tc := range testCases {
		t.Run(tc.method+" "+tc.path, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, server.URL+tc.path, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			// Verify status code
			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
			}

			// For non-HEAD requests, verify response body
			if tc.method != "HEAD" {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					t.Fatalf("Failed to read response body: %v", err)
				}

				expectedBody := "OK\n"
				if string(body) != expectedBody {
					t.Errorf("Expected body %q, got %q", expectedBody, string(body))
				}
			}
		})
	}
}