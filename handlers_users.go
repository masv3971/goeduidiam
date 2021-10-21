package goeduidiam

import (
	"context"
	"fmt"
	"net/http"
)

// UsersService holds the object for Users
type UsersService struct {
	client *Client
}

// UsersGetRequest type
type UsersGetRequest struct {
	ScimID string
}

// Get gets users from eduidIAM
func (s *UsersService) Get(ctx context.Context, request UsersGetRequest) (*UsersReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("/users/%s", request.ScimID),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &UsersReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// UsersPutRequest xx
type UsersPutRequest struct {
	ScimID string
	Data   UsersRequest `json:"data"`
}

// Put puts on user
func (s *UsersService) Put(ctx context.Context, request UsersPutRequest) (*UsersReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/users/%s", request.ScimID),
		request.Data,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &UsersReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// UsersPostRequest type
type UsersPostRequest struct {
	Data UsersRequest `json:"data"`
}

// Post posts user
func (s *UsersService) Post(ctx context.Context, request UsersPostRequest) (*UsersReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"POST",
		fmt.Sprintf("/users/"),
		request.Data,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &UsersReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// UsersSearchRequest type
type UsersSearchRequest struct {
	Data SearchRequest `json:"data"`
}

// Search searching for user
func (s *UsersService) Search(ctx context.Context, request UsersSearchRequest) (*UsersReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"POST",
		fmt.Sprintf("/users/"),
		request.Data,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &UsersReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}
