package goeduidiam

import (
	"context"
	"fmt"
	"net/http"
)

// LoginService holds the service for login
type LoginService struct {
	client *Client
}

// LoginPostRequest type
type LoginPostRequest struct {
	Data LoginRequest `json:"data"`
}

// Post login a user
func (s *LoginService) Post(ctx context.Context, request LoginPostRequest) (*EmptyStruct, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"POST",
		fmt.Sprintf("/login/"),
		request.Data,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &EmptyStruct{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}
