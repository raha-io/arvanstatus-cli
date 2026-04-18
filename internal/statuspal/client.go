// Package statuspal is a thin client for the StatusPal public API, scoped to
// the endpoints the arvanstatus-cli needs. All exported methods are safe to
// call from multiple goroutines.
package statuspal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	defaultBaseURL = "https://statuspal.io/api/v2"
	defaultPage    = "arvancloud"
	userAgent      = "arvanstatus-cli/0.1"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	page       string
}

type Option func(*Client)

func WithHTTPClient(h *http.Client) Option { return func(c *Client) { c.httpClient = h } }
func WithBaseURL(u string) Option          { return func(c *Client) { c.baseURL = u } }
func WithPage(p string) Option             { return func(c *Client) { c.page = p } }

func NewClient(opts ...Option) *Client {
	c := &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		baseURL:    defaultBaseURL,
		page:       defaultPage,
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

func (c *Client) getJSON(ctx context.Context, path string, out any) error {
	url := fmt.Sprintf("%s/status_pages/%s/%s", c.baseURL, c.page, path)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("statuspal: build request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("statuspal: GET %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return fmt.Errorf("statuspal: GET %s: status %d: %s", url, resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("statuspal: decode %s: %w", url, err)
	}
	return nil
}
