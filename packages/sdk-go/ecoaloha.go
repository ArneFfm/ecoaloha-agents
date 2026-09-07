// Package ecoaloha provides a client for the public, keyless EcoAloha API.
package ecoaloha

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const DefaultBaseURL = "https://ecoaloha.com/api/v1"

// Client calls the public API. Use NewClient to create one.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient uses the production API and http.DefaultClient for empty arguments.
func NewClient(baseURL string, httpClient *http.Client) (*Client, error) {
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	u, err := url.Parse(baseURL)
	if err != nil || u.Host == "" || (u.Scheme != "http" && u.Scheme != "https") || u.User != nil || u.RawQuery != "" || u.Fragment != "" {
		return nil, fmt.Errorf("use an HTTP(S) base URL without credentials, query, or fragment")
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{strings.TrimRight(baseURL, "/"), httpClient}, nil
}

// HTTPError preserves the response body and the Retry-After header.
type HTTPError struct {
	StatusCode int
	Body       json.RawMessage
	RetryAfter string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("EcoAloha request failed (%d)", e.StatusCode)
}

// RequestOptions describes a documented API operation. Method defaults to GET.
// Body is JSON-encoded when non-nil. Query supports repeated parameter values.
type RequestOptions struct {
	Method string
	Query  url.Values
	Body   any
}

var relativePath = regexp.MustCompile(`^/[a-zA-Z0-9_/-]*$`)

// Request returns the complete JSON response, including pagination metadata.
// It performs no automatic retries.
func (c *Client) Request(ctx context.Context, path string, options RequestOptions) (json.RawMessage, error) {
	if !relativePath.MatchString(path) || strings.Contains(path, "//") {
		return nil, fmt.Errorf("use an API-relative path, such as /destinations")
	}
	endpoint := c.baseURL + path
	if query := options.Query.Encode(); query != "" {
		endpoint += "?" + query
	}
	var body io.Reader
	if options.Body != nil {
		encoded, err := json.Marshal(options.Body)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(encoded)
	}
	method := options.Method
	if method == "" {
		method = http.MethodGet
	}
	req, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if options.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	payload, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, &HTTPError{response.StatusCode, payload, response.Header.Get("Retry-After")}
	}
	if !json.Valid(payload) {
		return nil, fmt.Errorf("EcoAloha returned invalid JSON")
	}
	return json.RawMessage(payload), nil
}

func (c *Client) Destinations(ctx context.Context) (json.RawMessage, error) {
	return c.Request(ctx, "/destinations", RequestOptions{})
}

func (c *Client) Experiences(ctx context.Context, query url.Values) (json.RawMessage, error) {
	return c.Request(ctx, "/experiences", RequestOptions{Query: query})
}

// Compare uses EUR when currency is empty.
func (c *Client) Compare(ctx context.Context, experienceIDs []string, currency string) (json.RawMessage, error) {
	if currency == "" {
		currency = "EUR"
	}
	return c.Request(ctx, "/compare", RequestOptions{
		Method: http.MethodPost,
		Body: struct {
			ExperienceIDs []string `json:"experienceIds"`
			Currency      string   `json:"currency"`
		}{experienceIDs, currency},
	})
}
