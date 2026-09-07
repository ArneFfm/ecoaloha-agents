package ecoaloha

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestOperations(t *testing.T) {
	const envelope = `{"data":[],"nextCursor":"next","extra":true}`
	tests := []struct {
		name, method, path, query, body string
		call                            func(*Client) (json.RawMessage, error)
	}{
		{"destinations", "GET", "/api/sandbox/v1/destinations", "", "", func(c *Client) (json.RawMessage, error) { return c.Destinations(context.Background()) }},
		{"experiences", "GET", "/api/sandbox/v1/experiences", "destinationId=paris&search=food+%26+wine", "", func(c *Client) (json.RawMessage, error) {
			return c.Experiences(context.Background(), url.Values{"destinationId": {"paris"}, "search": {"food & wine"}})
		}},
		{"compare", "POST", "/api/sandbox/v1/compare", "", `{"experienceIds":["fictional-1","fictional-2"],"currency":"EUR"}`, func(c *Client) (json.RawMessage, error) {
			return c.Compare(context.Background(), []string{"fictional-1", "fictional-2"}, "")
		}},
		{"custom", "POST", "/api/sandbox/v1/compare", "tag=a&tag=b", `{"currency":"USD"}`, func(c *Client) (json.RawMessage, error) {
			return c.Request(context.Background(), "/compare", RequestOptions{Method: "POST", Query: url.Values{"tag": {"a", "b"}}, Body: map[string]string{"currency": "USD"}})
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				body, _ := io.ReadAll(r.Body)
				if r.Method != tt.method || r.URL.Path != tt.path || r.URL.RawQuery != tt.query || string(body) != tt.body {
					t.Errorf("unexpected request: %s %s body=%s", r.Method, r.URL, body)
				}
				if r.Header.Get("Accept") != "application/json" || r.Header.Get("Authorization") != "" {
					t.Errorf("unexpected headers: %v", r.Header)
				}
				if tt.body != "" && r.Header.Get("Content-Type") != "application/json" {
					t.Error("missing JSON content type")
				}
				io.WriteString(w, envelope)
			}))
			defer server.Close()
			client, err := NewClient(server.URL+"/api/sandbox/v1/", server.Client())
			if err != nil {
				t.Fatal(err)
			}
			result, err := tt.call(client)
			if err != nil || string(result) != envelope {
				t.Fatalf("response=%s error=%v", result, err)
			}
		})
	}
}

func TestErrorAndCancellation(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.Header().Set("Retry-After", "30")
		w.WriteHeader(http.StatusTooManyRequests)
		io.WriteString(w, `{"detail":"Slow down"}`)
	}))
	defer server.Close()
	client, err := NewClient(server.URL, server.Client())
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Destinations(context.Background())
	var apiError *HTTPError
	if !errors.As(err, &apiError) || apiError.StatusCode != 429 || apiError.RetryAfter != "30" || string(apiError.Body) != `{"detail":"Slow down"}` {
		t.Fatalf("unexpected error: %#v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := client.Destinations(ctx); !errors.Is(err, context.Canceled) {
		t.Fatalf("expected cancellation, got %v", err)
	}
	if calls != 1 {
		t.Fatalf("expected one request without retries, got %d", calls)
	}
}

func TestRejectUnsafePaths(t *testing.T) {
	client, err := NewClient("", nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{"https://example.com", "//example.com", "/../admin", "/%2e%2e/admin", "/destinations?x=1", "/destinations#x", "/a//b", "/a\\b", "destinations"} {
		if _, err := client.Request(context.Background(), path, RequestOptions{}); err == nil {
			t.Errorf("accepted unsafe path %q", path)
		}
	}
	for _, base := range []string{"ftp://example.com", "https://user:pass@example.com", "https://example.com?x=1", "https://example.com#x", "/api/v1"} {
		if _, err := NewClient(base, nil); err == nil {
			t.Errorf("accepted invalid base URL %q", base)
		}
	}
}

func TestNonJSONResponses(t *testing.T) {
	for _, status := range []int{http.StatusOK, http.StatusBadGateway} {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(status)
			io.WriteString(w, "upstream unavailable")
		}))
		client, err := NewClient(server.URL, server.Client())
		if err != nil {
			t.Fatal(err)
		}
		_, err = client.Destinations(context.Background())
		server.Close()
		if err == nil {
			t.Fatal("expected an error for non-JSON response")
		}
		if status == http.StatusBadGateway {
			var apiError *HTTPError
			if !errors.As(err, &apiError) || string(apiError.Body) != "upstream unavailable" {
				t.Fatalf("expected preserved HTTP error body, got %v", err)
			}
		}
	}
}
