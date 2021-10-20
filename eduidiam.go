package goeduidiam

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Client holds the object
type Client struct {
	httpClient *http.Client
	URL        string

	Users *UsersService
}

// Config holds the configuration for New
type Config struct {
	URL string
}

// New creates a new instance of goeduidiam
func New(config Config) (*Client, error) {
	c := &Client{
		URL: config.URL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	c.Users = &UsersService{client: c}

	return c, nil
}

func (c *Client) newRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(c.URL)
	if err != nil {
		return nil, err
	}
	url := u.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		payload := struct {
			Data interface{} `json:"data"`
		}{
			Data: body,
		}
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(payload)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "goeduidiam/0.0.1")

	return req, nil
}

func (c *Client) do(req *http.Request, value interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := checkResponse(resp); err != nil {
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(value); err != nil {
		return nil, err
	}

	return resp, nil
}

func checkResponse(r *http.Response) error {
	serviceName := "goeduidiam"

	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	case 500:
		return fmt.Errorf("%s: not allowed", serviceName)
	}
	return fmt.Errorf("%s: invalid request", serviceName)
}
