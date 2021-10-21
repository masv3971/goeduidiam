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

// GetAll return all groups
func (s *GroupsService) GetAll(ctx context.Context) (*GroupsReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("/groups/"),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &GroupsReply{}
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

// Get gets one group
func (s *GroupsService) Get(ctx context.Context, request GroupsGetRequest) (*GroupsReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("/groups/%s", request.ScimID),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &GroupsReply{}
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

// Post posts a group
func (s *GroupsService) Post(ctx context.Context, request GroupsPostRequest) (*GroupsReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"POST",
		fmt.Sprintf("/groups/"),
		request.Data,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &GroupsReply{}
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

// Put puts group
func (s *GroupsService) Put(ctx context.Context, request GroupsPutRequest) (*GroupsReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/groups/%s", request.ScimID),
		request.Data,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &GroupsReply{}
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

// Delete deletes a group
func (s *GroupsService) Delete(ctx context.Context, request GroupsDeleteRequest) (*EmptyStruct, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/groups/%s", request.ScimID),
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

// GroupsSearchRequest is the request for Search
type GroupsSearchRequest struct {
	Data SearchRequest
}

// Search group
func (s *GroupsService) Search(ctx context.Context, request GroupsSearchRequest) (*SearchReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"POST",
		fmt.Sprintf("/groups/"),
		request.Data,
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
