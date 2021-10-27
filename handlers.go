package goeduidiam

import (
	"context"
	"net/http"
)

// UsersService holds the object for Users
type UsersService struct {
	client *Client
	path   string
}

// LoginService holds the service for login
type LoginService struct {
	client *Client
	path   string
}

// StatusService holds the object for Status
type StatusService struct {
	client *Client
	path   string
}

// InvitesService holds the object for Invites
type InvitesService struct {
	client *Client
	path   string
}

// GroupsService holds the service for groups
type GroupsService struct {
	client *Client
	path   string
}

// EventsService holds events object
type EventsService struct {
	client *Client
	path   string
}

// Get gets users from eduidIAM
func (s *UsersService) Get(ctx context.Context, req RequestCFG) (*UsersReply, *http.Response, error) {
	reply := &UsersReply{}
	resp, err := s.client.call(ctx, "GET", s.path, req.ScimID, nil, reply)
	if err != nil {
		return nil, nil, err
	}
	return reply, resp, nil
}

// Put puts on user
func (s *UsersService) Put(ctx context.Context, req RequestCFG) (*UsersReply, *http.Response, error) {
	reply := &UsersReply{}
	resp, err := s.client.call(ctx, "PUT", s.path, req.ScimID, req.Users, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// Post posts user
func (s *UsersService) Post(ctx context.Context, req RequestCFG) (*UsersReply, *http.Response, error) {
	reply := &UsersReply{}
	resp, err := s.client.call(ctx, "POST", s.path, "", req.Users, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// Search searching for user
func (s *UsersService) Search(ctx context.Context, req RequestCFG) (*SearchReply, *http.Response, error) {
	reply := &SearchReply{}
	resp, err := s.client.call(ctx, "POST", s.path, "", req.Search, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// GetHealthy gets health status
func (s *StatusService) GetHealthy(ctx context.Context) (*HealthyReply, *http.Response, error) {
	reply := &HealthyReply{}
	resp, err := s.client.call(ctx, "GET", "status/healthy", "", nil, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// Get gets one invite
func (s *InvitesService) Get(ctx context.Context, req RequestCFG) (*UsersReply, *http.Response, error) {
	reply := &UsersReply{}
	resp, err := s.client.call(ctx, "GET", s.path, req.ScimID, nil, reply)
	if err != nil {
		return nil, nil, err
	}
	return reply, resp, nil
}

// Delete deletes one invite
func (s *InvitesService) Delete(ctx context.Context, req RequestCFG) (*EmptyStruct, *http.Response, error) {
	reply := &EmptyStruct{}
	resp, err := s.client.call(ctx, "DELETE", s.path, req.ScimID, nil, reply)
	if err != nil {
		return nil, nil, err
	}
	return reply, resp, nil
}

// Post creates a new invite
func (s *InvitesService) Post(ctx context.Context, req RequestCFG) (*UsersReply, *http.Response, error) {
	reply := &UsersReply{}
	resp, err := s.client.call(ctx, "POST", s.path, "", req.Invites, reply)
	if err != nil {
		return nil, nil, err
	}
	return reply, resp, nil
}

// Search searches for an invite
func (s *InvitesService) Search(ctx context.Context, req RequestCFG) (*SearchReply, *http.Response, error) {
	reply := &SearchReply{}
	resp, err := s.client.call(ctx, "POST", s.path, "", req, reply)
	if err != nil {
		return nil, nil, err
	}
	return reply, resp, nil
}

// GetAll return all groups
func (s *GroupsService) GetAll(ctx context.Context, req RequestCFG) (*GroupsReply, *http.Response, error) {
	reply := &GroupsReply{}
	resp, err := s.client.call(ctx, "GET", s.path, "", nil, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, err
}

// Get gets one group
func (s *GroupsService) Get(ctx context.Context, req RequestCFG) (*GroupsReply, *http.Response, error) {
	reply := &GroupsReply{}
	resp, err := s.client.call(ctx, "GET", s.path, req.ScimID, nil, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// Post posts a group
func (s *GroupsService) Post(ctx context.Context, req RequestCFG) (*GroupsReply, *http.Response, error) {
	reply := &GroupsReply{}
	resp, err := s.client.call(ctx, "POST", s.path, "", req.Groups, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// Put puts group
func (s *GroupsService) Put(ctx context.Context, req RequestCFG) (*GroupsReply, *http.Response, error) {
	reply := &GroupsReply{}
	resp, err := s.client.call(ctx, "PUT", s.path, req.ScimID, req.Groups, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// Delete deletes a group
func (s *GroupsService) Delete(ctx context.Context, req RequestCFG) (*EmptyStruct, *http.Response, error) {
	reply := &EmptyStruct{}
	resp, err := s.client.call(ctx, "DELETE", s.path, req.ScimID, req.Groups, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// Search group
func (s *GroupsService) Search(ctx context.Context, req RequestCFG) (*SearchReply, *http.Response, error) {
	reply := &SearchReply{}
	resp, err := s.client.call(ctx, "POST", s.path, "", req.Search, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// Get gets one event
func (s *EventsService) Get(ctx context.Context, req RequestCFG) (*EventsReply, *http.Response, error) {
	reply := &EventsReply{}
	resp, err := s.client.call(ctx, "GET", s.path, req.ScimID, req.Events, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// Post create a new event
func (s *EventsService) Post(ctx context.Context, req RequestCFG) (*EventsReply, *http.Response, error) {
	reply := &EventsReply{}
	resp, err := s.client.call(ctx, "POST", s.path, "", req.Events, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}
