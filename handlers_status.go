package goeduidiam

import (
	"context"
	"fmt"
	"net/http"
)

// StatusService holds the object for Status
type StatusService struct {
	client *Client
}

// GetHealthy gets health status
func (s *StatusService) GetHealthy(ctx context.Context) (*HealthyReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("/status/healthy/"),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &HealthyReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}
