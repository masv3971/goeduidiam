package goeduidiam

import (
	"context"
	"fmt"
	"net/http"
)

// EventsService holds events object
type EventsService struct {
	client *Client
}

// EventsGetRequest type
type EventsGetRequest struct {
	ScimID string
}

// EventsGetReply type
type EventsGetReply struct {
	Ok  EventsReply
	Err ErrorReply
}

// Get gets one event
func (s *EventsService) Get(ctx context.Context, request EventsGetRequest) (*EventsGetReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("/events/%s", request.ScimID),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &EventsGetReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// EventsPostRequest type
type EventsPostRequest struct {
	Data EventsRequest
}

// EventsPostReply type
type EventsPostReply struct {
	Ok  EventsReply
	Err ErrorReply
}

// Post creats a new event
func (s *EventsService) Post(ctx context.Context, request EventsPostRequest) (*EventsPostReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"POST",
		fmt.Sprintf("/events/"),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &EventsPostReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}
