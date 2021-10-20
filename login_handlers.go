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
	Data struct {
		DataOwner string `json:"data_owner"`
	}
}

// LoginPostReply type
type LoginPostReply struct {
	Ok  struct{}
	Err ErrorReply
}

// Post login a user
func (s *LoginService) Post(ctx context.Context, request LoginPostRequest) (*LoginPostReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"POST",
		fmt.Sprintf("/users/"),
		request.Data,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &LoginPostReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}
