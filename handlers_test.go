package goeduidiam

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/masv3971/goeduidiam/eduidiammock"
	"github.com/stretchr/testify/assert"
)

func TestHandlers(t *testing.T) {
	var (
		client = mockNew(t, "")
	)

	tts := []struct {
		name             string
		serverMethod     string
		serverURL        string
		serverPayload    []byte
		serverStatusCode int
		clientReply      interface{}
		clientRequest    interface{}
		clientFn         interface{}
	}{
		{
			name:             "GetHealthy",
			serverMethod:     "GET",
			serverURL:        "/status/healthy/",
			serverPayload:    jsonHealthyReply,
			serverStatusCode: 200,
			clientReply:      &HealthyReply{},
			clientFn:         client.Status.GetHealthy,
		},
		{
			name:             "GetHealthy",
			serverMethod:     "GET",
			serverURL:        "/status/healthy/",
			serverPayload:    jsonErrorReply,
			serverStatusCode: 500,
			clientReply:      &Errors{},
			clientFn:         client.Status.GetHealthy,
		},
		{
			name:             "Get",
			serverMethod:     "GET",
			serverURL:        "/users/testuid",
			serverPayload:    jsonUsersReply,
			serverStatusCode: 200,
			clientReply:      &UsersReply{},
			clientRequest: &GetUsersRequest{
				ScimID: "testuid",
			},
			clientFn: client.Users.Get,
		},
		{
			name:             "Get",
			serverMethod:     "GET",
			serverURL:        "/users/testuid",
			serverPayload:    jsonErrorReply,
			serverStatusCode: 500,
			clientReply:      &Errors{},
			clientRequest: &GetUsersRequest{
				ScimID: "testuid",
			},
			clientFn: client.Users.Get,
		},
		{
			name:             "Put",
			serverMethod:     "PUT",
			serverURL:        "/users/testScim",
			serverPayload:    jsonUsersReply,
			serverStatusCode: 200,
			clientReply:      &UsersReply{},
			clientRequest:    &PutUsersRequest{ScimID: "testScim"},
			clientFn:         client.Users.Put,
		},
		{
			name:             "Put",
			serverMethod:     "PUT",
			serverURL:        "/users/testScim",
			serverPayload:    jsonErrorReply,
			serverStatusCode: 500,
			clientReply:      &Errors{},
			clientRequest:    &PutUsersRequest{ScimID: "testScim"},
			clientFn:         client.Users.Put,
		},
		{
			name:             "search",
			serverMethod:     "POST",
			serverURL:        "/users/.search/",
			serverPayload:    jsonSearchReply,
			serverStatusCode: 200,
			clientReply:      &SearchReply{},
			clientRequest:    &SearchUsersRequest{},
			clientFn:         client.Users.Search,
		},
		{
			name:             "search",
			serverMethod:     "POST",
			serverURL:        "/users/.search/",
			serverPayload:    jsonErrorReply,
			serverStatusCode: 500,
			clientReply:      &Errors{},
			clientRequest:    &SearchUsersRequest{},
			clientFn:         client.Users.Search,
		},
		{
			name:             "post",
			serverMethod:     "POST",
			serverURL:        "/users/",
			serverPayload:    jsonUsersReply,
			serverStatusCode: 200,
			clientReply:      &UsersReply{},
			clientRequest:    &PostUsersRequest{},
			clientFn:         client.Users.Post,
		},
		{
			name:             "post",
			serverMethod:     "POST",
			serverURL:        "/users/",
			serverPayload:    jsonErrorReply,
			serverStatusCode: 500,
			clientReply:      &Errors{},
			clientRequest:    &PostUsersRequest{},
			clientFn:         client.Users.Post,
		},
		{
			name:             "Get",
			serverMethod:     "GET",
			serverURL:        "/invites/testScim",
			serverPayload:    jsonUsersReply,
			serverStatusCode: 200,
			clientReply:      &UsersReply{},
			clientRequest:    &GetInvitesRequest{ScimID: "testScim"},
			clientFn:         client.Invites.Get,
		},
		{
			name:             "Get",
			serverMethod:     "GET",
			serverURL:        "/invites/testScim",
			serverPayload:    jsonErrorReply,
			serverStatusCode: 500,
			clientReply:      &Errors{},
			clientRequest:    &GetInvitesRequest{ScimID: "testScim"},
			clientFn:         client.Invites.Get,
		},
		{
			name:             "Post",
			serverMethod:     "POST",
			serverURL:        "/invites/",
			serverPayload:    jsonUsersReply,
			serverStatusCode: 200,
			clientReply:      &UsersReply{},
			clientRequest:    &PostInvitesRequest{},
			clientFn:         client.Invites.Post,
		},
		{
			name:             "Post",
			serverMethod:     "POST",
			serverURL:        "/invites/",
			serverPayload:    jsonErrorReply,
			serverStatusCode: 500,
			clientReply:      &Errors{},
			clientRequest:    &PostInvitesRequest{},
			clientFn:         client.Invites.Post,
		},
		{
			name:             "Search",
			serverMethod:     "POST",
			serverURL:        "/invites/.search/",
			serverPayload:    jsonUsersReply,
			serverStatusCode: 200,
			clientReply:      &SearchReply{},
			clientRequest:    &SearchInvitesRequest{},
			clientFn:         client.Invites.Search,
		},
		{
			name:             "Search",
			serverMethod:     "POST",
			serverURL:        "/invites/.search/",
			serverPayload:    jsonErrorReply,
			serverStatusCode: 500,
			clientReply:      &Errors{},
			clientRequest:    &SearchInvitesRequest{},
			clientFn:         client.Invites.Search,
		},
		{
			name:             "Delete",
			serverMethod:     "DELETE",
			serverURL:        "/invites/testScim",
			serverPayload:    jsonEmpty,
			serverStatusCode: 200,
			clientReply:      &EmptyStruct{},
			clientRequest:    &DeleteInvitesRequest{ScimID: "testScim"},
			clientFn:         client.Invites.Delete,
		},
		{
			name:             "Delete",
			serverMethod:     "DELETE",
			serverURL:        "/invites/testScim",
			serverPayload:    jsonErrorReply,
			serverStatusCode: 500,
			clientReply:      &Errors{},
			clientRequest:    &DeleteInvitesRequest{ScimID: "testScim"},
			clientFn:         client.Invites.Delete,
		},
		{
			name:             "Get",
			serverMethod:     "GET",
			serverURL:        "/groups/testScim",
			serverPayload:    jsonGroupsReply,
			serverStatusCode: 200,
			clientReply:      &GroupsReply{},
			clientRequest:    &GetGroupsRequest{ScimID: "testScim"},
			clientFn:         client.Groups.Get,
		},
		{
			name:             "Get",
			serverMethod:     "GET",
			serverURL:        "/groups/testScim",
			serverPayload:    jsonErrorReply,
			serverStatusCode: 500,
			clientReply:      &Errors{},
			clientRequest:    &GetGroupsRequest{ScimID: "testScim"},
			clientFn:         client.Groups.Get,
		},
		{
			name:             "GetAll",
			serverMethod:     "GET",
			serverURL:        "/groups/",
			serverPayload:    jsonGroupsReplyAll,
			serverStatusCode: 200,
			clientReply:      &GroupsReply{},
			clientFn:         client.Groups.GetAll,
		},
		{
			name:             "GetAll",
			serverMethod:     "GET",
			serverURL:        "/groups/",
			serverPayload:    jsonErrorReply,
			serverStatusCode: 500,
			clientReply:      &Errors{},
			clientFn:         client.Groups.GetAll,
		},
		{
			name:             "Post",
			serverMethod:     "POST",
			serverURL:        "/groups/",
			serverPayload:    jsonGroupsReply,
			serverStatusCode: 200,
			clientReply:      &GroupsReply{},
			clientRequest:    &PostGroupsRequest{},
			clientFn:         client.Groups.Post,
		},
		{
			name:             "Post",
			serverMethod:     "POST",
			serverURL:        "/groups/",
			serverPayload:    jsonErrorReply,
			serverStatusCode: 500,
			clientReply:      &Errors{},
			clientRequest:    &PostGroupsRequest{},
			clientFn:         client.Groups.Post,
		},
		{
			name:             "Put",
			serverMethod:     "PUT",
			serverURL:        "/groups/testScim",
			serverPayload:    jsonGroupsReply,
			serverStatusCode: 200,
			clientReply:      &GroupsReply{},
			clientRequest:    &PutGroupsRequest{ScimID: "testScim"},
			clientFn:         client.Groups.Put,
		},
		{
			name:             "Put",
			serverMethod:     "PUT",
			serverURL:        "/groups/testScim",
			serverPayload:    jsonErrorReply,
			serverStatusCode: 500,
			clientReply:      &Errors{},
			clientRequest:    &PutGroupsRequest{ScimID: "testScim"},
			clientFn:         client.Groups.Put,
		},
		{
			name:             "Search",
			serverMethod:     "POST",
			serverURL:        "/groups/.search/",
			serverPayload:    jsonSearchReply,
			serverStatusCode: 200,
			clientReply:      &SearchReply{},
			clientRequest:    &SearchGroupsRequest{},
			clientFn:         client.Groups.Search,
		},
		{
			name:             "Search",
			serverMethod:     "POST",
			serverURL:        "/groups/.search/",
			serverPayload:    jsonErrorReply,
			serverStatusCode: 500,
			clientReply:      &Errors{},
			clientRequest:    &SearchGroupsRequest{},
			clientFn:         client.Groups.Search,
		},
		{
			name:             "Delete",
			serverMethod:     "DELETE",
			serverURL:        "/groups/testScim",
			serverPayload:    jsonEmpty,
			serverStatusCode: 200,
			clientReply:      &EmptyStruct{},
			clientRequest:    &DeleteGroupsRequest{ScimID: "testScim"},
			clientFn:         client.Groups.Delete,
		},
		{
			name:             "Delete",
			serverMethod:     "DELETE",
			serverURL:        "/groups/testScim",
			serverPayload:    jsonErrorReply,
			serverStatusCode: 500,
			clientReply:      &Errors{},
			clientRequest:    &DeleteGroupsRequest{ScimID: "testScim"},
			clientFn:         client.Groups.Delete,
		},
		{
			name:             "Get",
			serverMethod:     "GET",
			serverURL:        "/events/testScim",
			serverPayload:    jsonEventsReply,
			serverStatusCode: 200,
			clientReply:      &EventsReply{},
			clientRequest:    &GetEventRequest{ScimID: "testScim"},
			clientFn:         client.Events.Get,
		},
		{
			name:             "Get",
			serverMethod:     "GET",
			serverURL:        "/events/testScim",
			serverPayload:    jsonErrorReply,
			serverStatusCode: 500,
			clientReply:      &Errors{},
			clientRequest:    &GetEventRequest{ScimID: "testScim"},
			clientFn:         client.Events.Get,
		},
		{
			name:             "Post",
			serverMethod:     "POST",
			serverURL:        "/events/",
			serverPayload:    jsonEventsReply,
			serverStatusCode: 200,
			clientReply:      &EventsReply{},
			clientRequest:    &PostEventsRequest{},
			clientFn:         client.Events.Post,
		},
		{
			name:             "Post",
			serverMethod:     "POST",
			serverURL:        "/events/",
			serverPayload:    jsonErrorReply,
			serverStatusCode: 500,
			clientReply:      &Errors{},
			clientRequest:    &PostEventsRequest{},
			clientFn:         client.Events.Post,
		},
	}

	for _, tt := range tts {
		t.Run(fmt.Sprintf("%s:%s %d -- %s", tt.serverMethod, tt.serverURL, tt.serverStatusCode, tt.name), func(t *testing.T) {
			mux, server, _ := mockSetup(t)
			client.url = server.URL          // add server url to client
			client.SunetJWT.URL = server.URL // add server url to sunetJWT client
			defer server.Close()

			err := json.Unmarshal(tt.serverPayload, tt.clientReply)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			mockGenericEndpointServer(t, mux, "POST", "/transaction", eduidiammock.MockJWTJSON, 200)                // JWT Endpoint
			mockGenericEndpointServer(t, mux, tt.serverMethod, tt.serverURL, tt.serverPayload, tt.serverStatusCode) // eduidiam Endpoint

			switch tt.clientFn.(type) {
			case func(context.Context) (*HealthyReply, *http.Response, error):
				f := tt.clientFn.(func(context.Context) (*HealthyReply, *http.Response, error))
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO())
					if !assert.NoError(t, err) {
						t.FailNow()
					}
					assert.Equal(t, tt.clientReply, reply, "Should be equal")
				case 500:
					r := tt.clientReply.(*Errors)

					_, _, err = f(context.TODO())
					assert.Equal(t, err.Error(), r.Error())
				}
			case func(context.Context, *GetUsersRequest) (*UsersReply, *http.Response, error):
				f := tt.clientFn.(func(context.Context, *GetUsersRequest) (*UsersReply, *http.Response, error))
				req := tt.clientRequest.(*GetUsersRequest)
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), req)
					if !assert.NoError(t, err) {
						t.FailNow()
					}
					assert.Equal(t, tt.clientReply, reply, "Should be equal")
				case 500:
					r := tt.clientReply.(*Errors)

					_, _, err = f(context.TODO(), req)
					assert.Equal(t, err.Error(), r.Error())
				}
			case func(context.Context, *PostUsersRequest) (*UsersReply, *http.Response, error):
				f := tt.clientFn.(func(context.Context, *PostUsersRequest) (*UsersReply, *http.Response, error))
				req := tt.clientRequest.(*PostUsersRequest)
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), req)
					if !assert.NoError(t, err) {
						t.FailNow()
					}
					assert.Equal(t, tt.clientReply, reply, "Should be equal")
				case 500:
					r := tt.clientReply.(*Errors)

					_, _, err = f(context.TODO(), req)
					assert.Equal(t, err.Error(), r.Error())
				}
			case func(context.Context, *SearchUsersRequest) (*SearchReply, *http.Response, error):
				f := tt.clientFn.(func(context.Context, *SearchUsersRequest) (*SearchReply, *http.Response, error))
				req := tt.clientRequest.(*SearchUsersRequest)
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), req)
					if !assert.NoError(t, err) {
						t.FailNow()
					}
					assert.Equal(t, tt.clientReply, reply, "Should be equal")
				case 500:
					r := tt.clientReply.(*Errors)

					_, _, err = f(context.TODO(), req)
					assert.Equal(t, err.Error(), r.Error())
				}
			case func(context.Context, *DeleteInvitesRequest) (*EmptyStruct, *http.Response, error):
				f := tt.clientFn.(func(context.Context, *DeleteInvitesRequest) (*EmptyStruct, *http.Response, error))
				req := tt.clientRequest.(*DeleteInvitesRequest)
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), req)
					if !assert.NoError(t, err) {
						t.FailNow()
					}
					assert.Equal(t, tt.clientReply, reply, "Should be equal")
				case 500:
					r := tt.clientReply.(*Errors)

					_, _, err = f(context.TODO(), req)
					assert.Equal(t, err.Error(), r.Error())
				}
			case func(context.Context, *DeleteGroupsRequest) (*EmptyStruct, *http.Response, error):
				req := tt.clientRequest.(*DeleteGroupsRequest)
				f := tt.clientFn.(func(context.Context, *DeleteGroupsRequest) (*EmptyStruct, *http.Response, error))
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), req)
					if !assert.NoError(t, err) {
						t.FailNow()
					}
					assert.Equal(t, tt.clientReply, reply, "Should be equal")
				case 500:
					r := tt.clientReply.(*Errors)

					_, _, err = f(context.TODO(), req)
					assert.Equal(t, err.Error(), r.Error())
				}
			case func(ctx context.Context, req *GetInvitesRequest) (*UsersReply, *http.Response, error):
				f := tt.clientFn.(func(ctx context.Context, req *GetInvitesRequest) (*UsersReply, *http.Response, error))
				req := tt.clientRequest.(*GetInvitesRequest)
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), req)
					if !assert.NoError(t, err) {
						t.FailNow()
					}
					assert.Equal(t, tt.clientReply, reply, "Should be equal")
				case 500:
					r := tt.clientReply.(*Errors)

					_, _, err = f(context.TODO(), req)
					assert.Equal(t, err.Error(), r.Error())
				}
			case func(context.Context, *PutUsersRequest) (*UsersReply, *http.Response, error):
				f := tt.clientFn.(func(context.Context, *PutUsersRequest) (*UsersReply, *http.Response, error))
				req := tt.clientRequest.(*PutUsersRequest)
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), req)
					if !assert.NoError(t, err) {
						t.FailNow()
					}
					assert.Equal(t, tt.clientReply, reply, "Should be equal")
				case 500:
					r := tt.clientReply.(*Errors)

					_, _, err = f(context.TODO(), req)
					assert.Equal(t, err.Error(), r.Error())
				}
			case func(context.Context, *PostInvitesRequest) (*UsersReply, *http.Response, error):
				f := tt.clientFn.(func(context.Context, *PostInvitesRequest) (*UsersReply, *http.Response, error))
				req := tt.clientRequest.(*PostInvitesRequest)
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), req)
					if !assert.NoError(t, err) {
						t.FailNow()
					}
					assert.Equal(t, tt.clientReply, reply, "Should be equal")
				case 500:
					r := tt.clientReply.(*Errors)

					_, _, err = f(context.TODO(), req)
					assert.Equal(t, err.Error(), r.Error())
				}
			case func(context.Context, *SearchInvitesRequest) (*SearchReply, *http.Response, error):
				f := tt.clientFn.(func(context.Context, *SearchInvitesRequest) (*SearchReply, *http.Response, error))
				req := tt.clientRequest.(*SearchInvitesRequest)
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), req)
					if !assert.NoError(t, err) {
						t.FailNow()
					}
					assert.Equal(t, tt.clientReply, reply, "Should be equal")
				case 500:
					r := tt.clientReply.(*Errors)

					_, _, err = f(context.TODO(), req)
					assert.Equal(t, err.Error(), r.Error())
				}
			case func(context.Context, *GetGroupsRequest) (*GroupsReply, *http.Response, error):
				f := tt.clientFn.(func(context.Context, *GetGroupsRequest) (*GroupsReply, *http.Response, error))
				req := tt.clientRequest.(*GetGroupsRequest)
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), req)
					if !assert.NoError(t, err) {
						t.FailNow()
					}
					assert.Equal(t, tt.clientReply, reply, "Should be equal")
				case 500:
					r := tt.clientReply.(*Errors)

					_, _, err = f(context.TODO(), req)
					assert.Equal(t, err.Error(), r.Error())
				}
			case func(context.Context) (*GroupsReply, *http.Response, error):
				f := tt.clientFn.(func(context.Context) (*GroupsReply, *http.Response, error))
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO())
					if !assert.NoError(t, err) {
						t.FailNow()
					}
					assert.Equal(t, tt.clientReply, reply, "Should be equal")
				case 500:
					r := tt.clientReply.(*Errors)

					_, _, err = f(context.TODO())
					assert.Equal(t, err.Error(), r.Error())
				}
			case func(context.Context, *PostGroupsRequest) (*GroupsReply, *http.Response, error):
				f := tt.clientFn.(func(context.Context, *PostGroupsRequest) (*GroupsReply, *http.Response, error))
				req := tt.clientRequest.(*PostGroupsRequest)
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), req)
					if !assert.NoError(t, err) {
						t.FailNow()
					}
					assert.Equal(t, tt.clientReply, reply, "Should be equal")
				case 500:
					r := tt.clientReply.(*Errors)

					_, _, err = f(context.TODO(), req)
					assert.Equal(t, err.Error(), r.Error())
				}
			case func(context.Context, *PutGroupsRequest) (*GroupsReply, *http.Response, error):
				f := tt.clientFn.(func(context.Context, *PutGroupsRequest) (*GroupsReply, *http.Response, error))
				req := tt.clientRequest.(*PutGroupsRequest)
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), req)
					if !assert.NoError(t, err) {
						t.FailNow()
					}
					assert.Equal(t, tt.clientReply, reply, "Should be equal")
				case 500:
					r := tt.clientReply.(*Errors)

					_, _, err = f(context.TODO(), req)
					assert.Equal(t, err.Error(), r.Error())
				}
			case func(context.Context, *SearchGroupsRequest) (*SearchReply, *http.Response, error):
				f := tt.clientFn.(func(context.Context, *SearchGroupsRequest) (*SearchReply, *http.Response, error))
				req := tt.clientRequest.(*SearchGroupsRequest)
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), req)
					if !assert.NoError(t, err) {
						t.FailNow()
					}
					assert.Equal(t, tt.clientReply, reply, "Should be equal")
				case 500:
					r := tt.clientReply.(*Errors)

					_, _, err = f(context.TODO(), req)
					assert.Equal(t, err.Error(), r.Error())
				}
			case func(context.Context, *GetEventRequest) (*EventsReply, *http.Response, error):
				f := tt.clientFn.(func(context.Context, *GetEventRequest) (*EventsReply, *http.Response, error))
				req := tt.clientRequest.(*GetEventRequest)
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), req)
					if !assert.NoError(t, err) {
						t.FailNow()
					}
					assert.Equal(t, tt.clientReply, reply, "Should be equal")
				case 500:
					r := tt.clientReply.(*Errors)

					_, _, err = f(context.TODO(), req)
					assert.Equal(t, err.Error(), r.Error())
				}
			case func(context.Context, *PostEventsRequest) (*EventsReply, *http.Response, error):
				f := tt.clientFn.(func(context.Context, *PostEventsRequest) (*EventsReply, *http.Response, error))
				req := tt.clientRequest.(*PostEventsRequest)
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), req)
					if !assert.NoError(t, err) {
						t.FailNow()
					}
					assert.Equal(t, tt.clientReply, reply, "Should be equal")
				case 500:
					r := tt.clientReply.(*Errors)

					_, _, err = f(context.TODO(), req)
					assert.Equal(t, err.Error(), r.Error())
				}

			default:
				t.Errorf("Can't find any matching function signatures %T", tt.clientFn)
			}

			server.Close() // Kill http server after each run
		})
	}
}
