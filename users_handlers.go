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

// UsersGetReply type
type UsersGetReply struct {
	Ok  UsersReply
	Err ErrorReply
}

// Get gets users from eduidIAM
func (s *UsersService) Get(ctx context.Context, request UsersGetRequest) (*UsersGetReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("/users/%s", request.ScimID),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &UsersGetReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// UsersPutRequest xx
type UsersPutRequest struct {
	ScimID string
	Data   UsersRequest
}

// UsersPutReply type
type UsersPutReply struct {
	Ok  UsersReply
	Err ErrorReply
}

// Put puts on user
func (s *UsersService) Put(ctx context.Context, request UsersPutRequest) (*UsersPutReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/users/%s", request.ScimID),
		request.Data,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &UsersPutReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// UsersPostRequest type
type UsersPostRequest struct {
	ScimID string
	Data   UsersRequest
}

// UsersPostReply type
type UsersPostReply struct {
	Ok  UsersReply
	Err ErrorReply
}

// Post posts user
func (s *UsersService) Post(ctx context.Context, request UsersPostRequest) (*UsersPostReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"POST",
		fmt.Sprintf("/users/%s", request.ScimID),
		request.Data,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &UsersPostReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// UsersSearchRequest type
type UsersSearchRequest struct {
	Data SearchRequest
}

// UsersSearchReply type
type UsersSearchReply struct {
	Ok  UsersReply
	Err ErrorReply
}

// Search searching for user
func (s *UsersService) Search(ctx context.Context, request UsersSearchRequest) (*UsersSearchReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/users/"),
		request.Data,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &UsersSearchReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}
