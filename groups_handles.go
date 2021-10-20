package goeduidiam

import (
	"context"
	"fmt"
	"net/http"
)

// GroupsService holds the service for groups
type GroupsService struct {
	client *Client
}

// GroupsGetAllReply response from /groups/get
type GroupsGetAllReply struct {
	Schemas   []string `json:"schemas"`
	Resources []struct {
	} `json:"Resources"`
	TotalResults int `json:"totalResults"`
}

// GetAll return all groups
func (s *GroupsService) GetAll(ctx context.Context) (*GroupsGetAllReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("/groups/"),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &GroupsGetAllReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// GroupsGetRequest type
type GroupsGetRequest struct {
	ScimID string
}

// GroupsGetReply type
type GroupsGetReply struct {
	Ok  GroupsReply
	Err ErrorReply
}

// Get gets one group
func (s *GroupsService) Get(ctx context.Context, request GroupsGetRequest) (*GroupsGetReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("/groups/%s", request.ScimID),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &GroupsGetReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// GroupsPostRequest is the request for
type GroupsPostRequest struct {
	Data GroupsRequest
}

// GroupsPostReply type
type GroupsPostReply struct {
	Ok  GroupsReply
	Err ErrorReply
}

// Post posts a group
func (s *GroupsService) Post(ctx context.Context, request GroupsPostRequest) (*GroupsPostReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"POST",
		fmt.Sprintf("/groups/"),
		request.Data,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &GroupsPostReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// GroupsPutRequest is the request for
type GroupsPutRequest struct {
	ScimID string
	Data   GroupsRequest
}

// GroupsPutReply type
type GroupsPutReply struct {
	Ok  GroupsReply
	Err ErrorReply
}

// Put puts group
func (s *GroupsService) Put(ctx context.Context, request GroupsPutRequest, scimID string) (*GroupsPutReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/groups/%s", request.ScimID),
		request.Data,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &GroupsPutReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// GroupsDeleteRequest type
type GroupsDeleteRequest struct {
	ScimID string
}

// GroupsDeleteReply type
type GroupsDeleteReply struct {
	Ok  struct{}
	Err ErrorReply
}

// Delete delets a group
func (s *GroupsService) Delete(ctx context.Context, request GroupsDeleteRequest) (*GroupsDeleteReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/groups/%s", request.ScimID),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &GroupsDeleteReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// GroupsSearchRequest is the request for Search
type GroupsSearchRequest struct {
	Data SearchRequest
}

// GroupsSearchReply is the reply for Search
type GroupsSearchReply struct {
	Ok  SearchReply
	Err ErrorReply
}

// Search group
func (s *GroupsService) Search(ctx context.Context, request GroupsSearchRequest) (*GroupsSearchReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"POST",
		fmt.Sprintf("/groups/"),
		request.Data,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &GroupsSearchReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}
