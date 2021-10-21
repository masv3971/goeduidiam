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

// Get gets one invite
func (s *InvitesService) Get(ctx context.Context, request InvitesGetRequest) (*UsersReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("/invites/%s", request.ScimID),
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

// InvitesDeleteRequest type
type InvitesDeleteRequest struct {
	ScimID string
}

// Delete deletes one invite
func (s *InvitesService) Delete(ctx context.Context, request InvitesDeleteRequest) (*EmptyStruct, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/invites/%s", request.ScimID),
		nil,
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

// InvitesPostRequest type
type InvitesPostRequest struct {
	Data UsersRequest
}

// Post creates a new invite
func (s *InvitesService) Post(ctx context.Context, request InvitesPostRequest) (*UsersReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"POST",
		fmt.Sprintf("/invites/"),
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

// InvitesSearchRequest type
type InvitesSearchRequest struct {
	Data SearchRequest
}

// Search searches for an invite
func (s *InvitesService) Search(ctx context.Context, request InvitesSearchRequest) (*SearchReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"POST",
		fmt.Sprintf("/invites/"),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &SearchReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}
