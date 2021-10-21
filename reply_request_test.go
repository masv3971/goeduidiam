package goeduidiam

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReplyRequestStructs(t *testing.T) {
	tts := []struct {
		name string
		s    interface{}
		json []byte
	}{
		{
			name: "UsersReply",
			s:    &UsersReply{},
			json: jsonUsersReply,
		},
		{
			name: "UsersRequest",
			s:    &UsersRequest{},
			json: jsonUsersRequest,
		},
		{
			name: "ErrorReply",
			s:    &Errors{},
			json: jsonErrorReply,
		},
		{
			name: "EventReply",
			s:    &EventsReply{},
			json: jsonEventsReply,
		},
		{
			name: "EventRequest",
			s:    &EventsRequest{},
			json: jsonEventsRequest,
		},
		{
			name: "SearchReply",
			s:    &SearchReply{},
			json: jsonSearchReply,
		},
		{
			name: "SearchRequest",
			s:    &SearchRequest{},
			json: jsonSearchRequest,
		},
		{
			name: "HealthyReply",
			s:    &HealthyReply{},
			json: jsonHealthyReply,
		},
		{
			name: "GroupsReply",
			s:    &GroupsReply{},
			json: jsonGroupsReply,
		},
		{
			name: "GroupsRequest",
			s:    &GroupsRequest{},
			json: jsonGroupsRequest,
		},
	}

	for _, tt := range tts {
		t.Logf("Now running test for %q", tt.name)

		err := json.Unmarshal(tt.json, tt.s)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		got, err := json.Marshal(tt.s)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		require.JSONEq(t, string(tt.json), string(got))
	}
}
