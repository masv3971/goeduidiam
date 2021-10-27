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

	Events  *EventsService
	Groups  *GroupsService
	Invites *InvitesService
	Login   *LoginService
	Status  *StatusService
	Users   *UsersService
}

// Config holds the configuration for New
type Config struct {
	URL string
}

// New creates a new instance of goeduidiam
func New(config Config) *Client {
	c := &Client{
		URL: config.URL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	c.Events = &EventsService{client: c, path: "events"}
	c.Groups = &GroupsService{client: c, path: "groups"}
	c.Invites = &InvitesService{client: c, path: "invites"}
	c.Login = &LoginService{client: c}
	c.Status = &StatusService{client: c, path: "status"}
	c.Users = &UsersService{client: c, path: "users"}

	return c
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
		errorReply := &Errors{}
		buf := &bytes.Buffer{}
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(buf.Bytes(), errorReply); err != nil {
			return nil, err
		}
		return nil, errorReply
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
	default:
		return fmt.Errorf("%s: invalid request", serviceName)
	}
}

func (c *Client) call(ctx context.Context, verb, path, param string, req, value interface{}) (*http.Response, error) {
	request, err := c.newRequest(
		ctx,
		verb,
		fmt.Sprintf("/%s/%s", path, param),
		req,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(request, value)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
