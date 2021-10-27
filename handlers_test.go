package goeduidiam

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandlersDelete(t *testing.T) {
	var (
		client = mockNew(t, "")
	)

	tts := []struct {
		name       string
		path       string
		reply      interface{}
		statusCode int
		fn         func(context.Context, RequestCFG) (*EmptyStruct, *http.Response, error)
		payload    []byte
	}{
		{
			name:       "invites-200",
			path:       "/invites",
			reply:      &EmptyStruct{},
			statusCode: 200,
			fn:         client.Invites.Delete,
			payload:    jsonEmpty,
		},
		{
			name:       "invites-500",
			path:       "/invites",
			reply:      &Errors{},
			statusCode: 500,
			fn:         client.Invites.Delete,
			payload:    jsonErrorReply,
		},
		{
			name:       "groups-200",
			path:       "/groups",
			reply:      &EmptyStruct{},
			statusCode: 200,
			fn:         client.Groups.Delete,
			payload:    jsonEmpty,
		},
		{
			name:       "groups-500",
			path:       "/groups",
			reply:      &Errors{},
			statusCode: 500,
			fn:         client.Groups.Delete,
			payload:    jsonErrorReply,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mux, server, _ := mockSetup(t)
			client.URL = server.URL // add server url to client

			err := json.Unmarshal(tt.payload, tt.reply)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			mockGenericEndpointServer(t, mux, "DELETE", tt.path, tt.payload, RequestCFG{}, tt.statusCode)

			switch tt.statusCode {
			case 200:
				reply, _, err := tt.fn(context.TODO(), RequestCFG{})
				if !assert.NoError(t, err) {
					t.Fatal(err)
				}
				assert.Equal(t, tt.reply, reply, "Should be equal")
			case 500:
				r := tt.reply.(*Errors)

				_, _, err = tt.fn(context.TODO(), RequestCFG{})
				assert.Equal(t, err.Error(), r.Error())
			}

			server.Close() // Kill http server after each run
		})
	}
}
func TestHandlersSearch(t *testing.T) {
	var (
		client = mockNew(t, "")
	)

	tts := []struct {
		name       string
		path       string
		reply      interface{}
		statusCode int
		fn         func(context.Context, RequestCFG) (*SearchReply, *http.Response, error)
		payload    []byte
		req        RequestCFG
	}{
		{
			name:       "users-200",
			path:       "/users",
			reply:      &SearchReply{},
			statusCode: 200,
			fn:         client.Users.Search,
			payload:    jsonSearchReply,
			req:        RequestCFG{},
		},
		{
			name:       "users-500",
			path:       "/users",
			reply:      &Errors{},
			statusCode: 500,
			fn:         client.Users.Search,
			payload:    jsonErrorReply,
			req:        RequestCFG{},
		},
		{
			name:       "groups-200",
			path:       "/groups",
			reply:      &SearchReply{},
			statusCode: 200,
			fn:         client.Groups.Search,
			payload:    jsonSearchReply,
			req:        RequestCFG{},
		},
		{
			name:       "groups-500",
			path:       "/groups",
			reply:      &Errors{},
			statusCode: 500,
			fn:         client.Groups.Search,
			payload:    jsonErrorReply,
			req:        RequestCFG{},
		},
		{
			name:       "invites-200",
			path:       "/invites",
			reply:      &SearchReply{},
			statusCode: 200,
			fn:         client.Invites.Search,
			payload:    jsonSearchReply,
			req:        RequestCFG{},
		},
		{
			name:       "invites-500",
			path:       "/invites",
			reply:      &Errors{},
			statusCode: 500,
			fn:         client.Invites.Search,
			payload:    jsonErrorReply,
			req:        RequestCFG{},
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mux, server, _ := mockSetup(t)
			client.URL = server.URL // add server url to client

			err := json.Unmarshal(tt.payload, tt.reply)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			mockGenericEndpointServer(t, mux, "POST", tt.path, tt.payload, tt.req, tt.statusCode)

			switch tt.statusCode {
			case 200:
				reply, _, err := tt.fn(context.TODO(), tt.req)
				if !assert.NoError(t, err) {
					t.Fatal(err)
				}
				assert.Equal(t, tt.reply, reply, "Should be equal")
			case 500:
				r := tt.reply.(*Errors)

				_, _, err = tt.fn(context.TODO(), tt.req)
				assert.Equal(t, err.Error(), r.Error())
			}

			server.Close() // Kill http server after each run
		})
	}
}

func TestHandlersStatus(t *testing.T) {
	var (
		client = mockNew(t, "")
	)

	tts := []struct {
		name       string
		reply      interface{}
		verb       string
		path       string
		payload    []byte
		statusCode int
		fn         func(context.Context) (*HealthyReply, *http.Response, error)
	}{
		{
			name:       "GetHealthy 200",
			reply:      &HealthyReply{},
			verb:       http.MethodGet,
			path:       "/status/healthy",
			payload:    jsonHealthyReply,
			statusCode: 200,
			fn:         client.Status.GetHealthy,
		},
		{
			name:       "GetHealthy 500",
			reply:      &Errors{},
			verb:       http.MethodGet,
			path:       "/status/healthy",
			payload:    jsonErrorReply,
			statusCode: 500,
			fn:         client.Status.GetHealthy,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mux, server, _ := mockSetup(t)
			client.URL = server.URL // Add server url to *Client

			mockGenericEndpointServer(t, mux, tt.verb, tt.path, tt.payload, RequestCFG{}, tt.statusCode)

			err := json.Unmarshal(tt.payload, tt.reply)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			switch tt.statusCode {
			case 200:
				reply, _, err := tt.fn(context.TODO())
				if !assert.NoError(t, err) {
					t.Fatal(err)
				}

				if !assert.Equal(t, tt.reply, reply, "Should be equal") {
					t.FailNow()
				}
			case 500:
				r := tt.reply.(*Errors)

				_, _, err = tt.fn(context.TODO())
				assert.Equal(t, err.Error(), r.Error())
			}

			server.Close() // Close server after each run
		})
	}
}

func TestHandlersGroup(t *testing.T) {
	var (
		client = mockNew(t, "")
	)

	tts := []struct {
		name       string
		reply      interface{}
		req        RequestCFG
		verb       string
		path       string
		payload    []byte
		statusCode int
		fn         func(context.Context, RequestCFG) (*GroupsReply, *http.Response, error)
	}{
		{
			name:       "Get:/groups 200",
			reply:      &GroupsReply{},
			req:        RequestCFG{},
			verb:       http.MethodGet,
			path:       "/groups",
			payload:    jsonGroupsReply,
			statusCode: 200,
			fn:         client.Groups.Get,
		},
		{
			name:       "Get:/groups 500",
			reply:      &Errors{},
			req:        RequestCFG{},
			verb:       http.MethodGet,
			path:       "/groups",
			payload:    jsonErrorReply,
			statusCode: 500,
			fn:         client.Groups.Get,
		},
		{
			name:       "GetAll:/groups 200",
			reply:      &GroupsReply{},
			req:        RequestCFG{},
			verb:       http.MethodGet,
			path:       "/groups",
			payload:    jsonGroupsReply,
			statusCode: 200,
			fn:         client.Groups.GetAll,
		},
		{
			name:       "GetAll:/groups 500",
			reply:      &Errors{},
			req:        RequestCFG{},
			verb:       http.MethodGet,
			path:       "/groups",
			payload:    jsonErrorReply,
			statusCode: 500,
			fn:         client.Groups.GetAll,
		},
		{
			name:       "Post:/groups 200",
			reply:      &GroupsReply{},
			req:        RequestCFG{},
			verb:       http.MethodPost,
			path:       "/groups",
			payload:    jsonGroupsReply,
			statusCode: 200,
			fn:         client.Groups.Post,
		},
		{
			name:       "Post:/groups 500",
			reply:      &Errors{},
			req:        RequestCFG{},
			verb:       http.MethodPost,
			path:       "/groups",
			payload:    jsonErrorReply,
			statusCode: 500,
			fn:         client.Groups.Post,
		},
		{
			name:       "Put:/groups 200",
			reply:      &GroupsReply{},
			req:        RequestCFG{},
			verb:       http.MethodPut,
			path:       "/groups",
			payload:    jsonGroupsReply,
			statusCode: 200,
			fn:         client.Groups.Put,
		},
		{
			name:       "Put:/groups 500",
			reply:      &Errors{},
			req:        RequestCFG{},
			verb:       http.MethodPut,
			path:       "/groups",
			payload:    jsonErrorReply,
			statusCode: 500,
			fn:         client.Groups.Put,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mux, server, _ := mockSetup(t)
			client.URL = server.URL // Add server url to *Client

			mockGenericEndpointServer(t, mux, tt.verb, tt.path, tt.payload, tt.req, tt.statusCode)

			err := json.Unmarshal(tt.payload, tt.reply)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			switch tt.statusCode {
			case 200:
				reply, _, err := tt.fn(context.TODO(), tt.req)
				if !assert.NoError(t, err) {
					t.Fatal(err)
				}

				if !assert.Equal(t, tt.reply, reply, "Should be equal") {
					t.FailNow()
				}
			case 500:
				r := tt.reply.(*Errors)

				_, _, err = tt.fn(context.TODO(), tt.req)
				assert.Equal(t, err.Error(), r.Error())
			}

			server.Close() // Close server after each run
		})
	}
}

func TestHandlersUser(t *testing.T) {
	var (
		client = mockNew(t, "") // init *Client
	)

	tts := []struct {
		name       string
		reply      interface{}
		req        RequestCFG
		verb       string
		path       string
		payload    []byte
		statusCode int
		fn         func(context.Context, RequestCFG) (*UsersReply, *http.Response, error)
	}{
		{
			name:       "Get:/users 200",
			reply:      &UsersReply{},
			req:        RequestCFG{ScimID: uuid.NewString()},
			verb:       http.MethodGet,
			path:       "/users",
			fn:         client.Users.Get,
			payload:    jsonUsersReply,
			statusCode: 200,
		},
		{
			name:       "Get:/users 500",
			reply:      &Errors{},
			req:        RequestCFG{ScimID: uuid.NewString()},
			verb:       http.MethodGet,
			path:       "/users",
			fn:         client.Users.Get,
			payload:    jsonErrorReply,
			statusCode: 500,
		},
		{
			name:       "Post:/users 200",
			reply:      &UsersReply{},
			req:        RequestCFG{},
			verb:       http.MethodPost,
			path:       "/users",
			fn:         client.Users.Post,
			payload:    jsonUsersReply,
			statusCode: 200,
		},
		{
			name:       "Post:/users 500",
			reply:      &Errors{},
			req:        RequestCFG{},
			verb:       http.MethodPost,
			path:       "/users",
			fn:         client.Users.Post,
			payload:    jsonErrorReply,
			statusCode: 500,
		},
		{
			name:       "Put:/users 200",
			reply:      &UsersReply{},
			req:        RequestCFG{ScimID: uuid.NewString()},
			verb:       http.MethodPut,
			path:       "/users",
			fn:         client.Users.Put,
			payload:    jsonUsersReply,
			statusCode: 200,
		},
		{
			name:       "Put:/users 500",
			reply:      &Errors{},
			req:        RequestCFG{ScimID: uuid.NewString()},
			verb:       http.MethodPut,
			path:       "/users",
			fn:         client.Users.Put,
			payload:    jsonErrorReply,
			statusCode: 500,
		},
		{
			name:       "Post:/invites 200",
			reply:      &UsersReply{},
			req:        RequestCFG{},
			verb:       http.MethodPost,
			path:       "/invites",
			fn:         client.Invites.Post,
			payload:    jsonUsersReply,
			statusCode: 200,
		},
		{
			name:       "Post:/invites 500",
			reply:      &Errors{},
			req:        RequestCFG{},
			verb:       http.MethodPost,
			path:       "/invites",
			fn:         client.Invites.Post,
			payload:    jsonErrorReply,
			statusCode: 500,
		},
		{
			name:       "Get:/invites 200",
			reply:      &UsersReply{},
			req:        RequestCFG{ScimID: uuid.NewString()},
			verb:       http.MethodGet,
			path:       "/invites",
			fn:         client.Invites.Get,
			payload:    jsonUsersReply,
			statusCode: 200,
		},
		{
			name:       "Get:/invites 500",
			reply:      &Errors{},
			req:        RequestCFG{ScimID: uuid.NewString()},
			verb:       http.MethodGet,
			path:       "/invites",
			fn:         client.Invites.Get,
			payload:    jsonErrorReply,
			statusCode: 500,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mux, server, _ := mockSetup(t)
			client.URL = server.URL // Add server url to *Client

			mockGenericEndpointServer(t, mux, tt.verb, tt.path, tt.payload, tt.req, tt.statusCode)

			err := json.Unmarshal(tt.payload, tt.reply)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			switch tt.statusCode {
			case 200:
				reply, _, err := tt.fn(context.TODO(), tt.req)
				if !assert.NoError(t, err) {
					t.Fatal(err)
				}

				if !assert.Equal(t, tt.reply, reply, "Should be equal") {
					t.FailNow()
				}
			case 500:
				r := tt.reply.(*Errors)

				_, _, err = tt.fn(context.TODO(), tt.req)
				assert.Equal(t, err.Error(), r.Error())
			}

			server.Close() // Close server after each run
		})
	}
}

func mockGenericEndpointServer(t *testing.T, mux *http.ServeMux, verb, path string, payload []byte, req RequestCFG, statusCode int) {
	mux.HandleFunc(fmt.Sprintf("%s/%s", path, req.ScimID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(statusCode)
			testMethod(t, r, verb)
			testURL(t, r, fmt.Sprintf("%s/%s", path, req.ScimID))
			w.Write(payload)
		},
	)
}

func TestHandlersEvent(t *testing.T) {
	var (
		client = mockNew(t, "")
	)

	tts := []struct {
		name       string
		reply      interface{}
		req        RequestCFG
		verb       string
		path       string
		payload    []byte
		statusCode int
		fn         func(context.Context, RequestCFG) (*EventsReply, *http.Response, error)
	}{
		{
			name:       "Get:/events 200",
			reply:      &EventsReply{},
			req:        RequestCFG{},
			verb:       http.MethodGet,
			path:       "/events",
			payload:    jsonEventsReply,
			statusCode: 200,
			fn:         client.Events.Get,
		},
		{
			name:       "Get:/events 500",
			reply:      &Errors{},
			req:        RequestCFG{},
			verb:       http.MethodGet,
			path:       "/events",
			payload:    jsonErrorReply,
			statusCode: 500,
			fn:         client.Events.Get,
		},
		{
			name:       "Post:/events 200",
			reply:      &EventsReply{},
			req:        RequestCFG{},
			verb:       http.MethodPost,
			path:       "/events",
			payload:    jsonEventsReply,
			statusCode: 200,
			fn:         client.Events.Post,
		},
		{
			name:       "Post:/events 500",
			reply:      &Errors{},
			req:        RequestCFG{},
			verb:       http.MethodPost,
			path:       "/events",
			payload:    jsonErrorReply,
			statusCode: 500,
			fn:         client.Events.Post,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mux, server, _ := mockSetup(t)
			client.URL = server.URL // Add server url to *Client

			mockGenericEndpointServer(t, mux, tt.verb, tt.path, tt.payload, tt.req, tt.statusCode)

			err := json.Unmarshal(tt.payload, tt.reply)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			switch tt.statusCode {
			case 200:
				reply, _, err := tt.fn(context.TODO(), tt.req)
				if !assert.NoError(t, err) {
					t.Fatal(err)
				}

				if !assert.Equal(t, tt.reply, reply, "Should be equal") {
					t.FailNow()
				}
			case 500:
				r := tt.reply.(*Errors)

				_, _, err = tt.fn(context.TODO(), tt.req)
				assert.Equal(t, err.Error(), r.Error())
			}

			server.Close() // Close server after each run
		})
	}
}
