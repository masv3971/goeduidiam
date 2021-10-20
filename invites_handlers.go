package goeduidiam

import (
	"context"
	"fmt"
	"net/http"
)

// InvitesService holds the object for Invites
type InvitesService struct {
	client *Client
}

// InvitesGetRequest type
type InvitesGetRequest struct {
	ScimID string
}

// InvitesGetReply type
type InvitesGetReply struct {
	Ok  UsersReply
	Err ErrorReply
}

// Get gets one invite
func (s *InvitesService) Get(ctx context.Context, request InvitesGetRequest) (*InvitesGetReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("/invite/%s", request.ScimID),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &InvitesGetReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// InvitesDeleteRequest type
type InvitesDeleteRequest struct {
	ScimID string
}

// InvitesDeleteReply type
type InvitesDeleteReply struct {
	Ok struct{}
}

// Delete deletes one invite
func (s *InvitesService) Delete(ctx context.Context, request InvitesDeleteRequest) (*InvitesDeleteReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("/invite/%s", request.ScimID),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &InvitesDeleteReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// InvitesPostRequest type
type InvitesPostRequest struct {
	Data UsersRequest
}

// InvitesPostReply type
type InvitesPostReply struct {
	Ok  UsersReply
	Err ErrorReply
}

// Post creates a new invite
func (s *InvitesService) Post(ctx context.Context, request InvitesPostRequest) (*InvitesPostReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"POST",
		fmt.Sprintf("/invite/"),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &InvitesPostReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// InvitesSearchRequest type
type InvitesSearchRequest struct {
	Data SearchRequest
}

// InvitesSearchReply type
type InvitesSearchReply struct {
	Ok  SearchReply
	Err ErrorReply
}

// Search searches for an invite
func (s *InvitesService) Search(ctx context.Context, request InvitesPostRequest) (*InvitesPostReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"POST",
		fmt.Sprintf("/invite/"),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &InvitesPostReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}
