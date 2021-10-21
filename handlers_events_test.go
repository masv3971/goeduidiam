package goeduidiam

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventsGet_happy(t *testing.T) {
	d := &EventsReply{}

	err := json.Unmarshal(jsonEventsReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := EventsGetRequest{
		ScimID: "test",
	}

	mux.HandleFunc(fmt.Sprintf("/events/%s", req.ScimID),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/events/%s", req.ScimID))
			w.Write(jsonEventsReply)
		},
	)

	reply, _, err := client.Events.Get(context.TODO(), req)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestEventsGet_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := EventsGetRequest{
		ScimID: "test",
	}

	mux.HandleFunc(fmt.Sprintf("/events/%s", req.ScimID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/events/%s", req.ScimID))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Events.Get(context.TODO(), req)
	assert.Equal(t, err.Error(), d.Error())
}

func TestEventsPost_happy(t *testing.T) {
	d := &EventsReply{}

	err := json.Unmarshal(jsonEventsReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := EventsPostRequest{}

	mux.HandleFunc(fmt.Sprintf("/events/"),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			testURL(t, r, fmt.Sprintf("/events/"))
			w.Write(jsonEventsReply)
		},
	)

	reply, _, err := client.Events.Post(context.TODO(), req)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestEventsPost_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := EventsPostRequest{}

	mux.HandleFunc(fmt.Sprintf("/events/"),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "POST")
			testURL(t, r, fmt.Sprintf("/events/"))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Events.Post(context.TODO(), req)
	assert.Equal(t, err.Error(), d.Error())
}
