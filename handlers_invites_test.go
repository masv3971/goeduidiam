package goeduidiam

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvitesGet_happy(t *testing.T) {
	d := &UsersReply{}

	err := json.Unmarshal(jsonUsersReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := InvitesGetRequest{
		ScimID: "testID",
	}

	mux.HandleFunc(fmt.Sprintf("/invites/%s", req.ScimID),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/invites/%s", req.ScimID))
			w.Write(jsonUsersReply)
		},
	)

	reply, _, err := client.Invites.Get(context.TODO(), req)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestInvitesGet_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := InvitesGetRequest{
		ScimID: "testID",
	}

	mux.HandleFunc(fmt.Sprintf("/invites/%s", req.ScimID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/invites/%s", req.ScimID))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Invites.Get(context.TODO(), req)
	assert.Equal(t, err.Error(), d.Error())
}

func TestInvitesDelete_happy(t *testing.T) {
	d := &EmptyStruct{}

	err := json.Unmarshal(jsonEmpty, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := InvitesDeleteRequest{
		ScimID: "testID",
	}

	mux.HandleFunc(fmt.Sprintf("/invites/%s", req.ScimID),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
			testURL(t, r, fmt.Sprintf("/invites/%s", req.ScimID))
			w.Write(jsonEmpty)
		},
	)

	reply, _, err := client.Invites.Delete(context.TODO(), req)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestInvitesDelete_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := InvitesDeleteRequest{
		ScimID: "testID",
	}

	mux.HandleFunc(fmt.Sprintf("/invites/%s", req.ScimID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "DELETE")
			testURL(t, r, fmt.Sprintf("/invites/%s", req.ScimID))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Invites.Delete(context.TODO(), req)
	assert.Equal(t, err.Error(), d.Error())
}

func TestInvitesPost_happy(t *testing.T) {
	d := &UsersReply{}

	err := json.Unmarshal(jsonUsersReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := InvitesPostRequest{}

	mux.HandleFunc(fmt.Sprintf("/invites/"),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			testURL(t, r, fmt.Sprintf("/invites/"))
			w.Write(jsonUsersReply)
		},
	)

	reply, _, err := client.Invites.Post(context.TODO(), req)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestInvitesPost_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := InvitesPostRequest{}

	mux.HandleFunc(fmt.Sprintf("/invites/"),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "POST")
			testURL(t, r, fmt.Sprintf("/invites/"))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Invites.Post(context.TODO(), req)
	assert.Equal(t, err.Error(), d.Error())
}

func TestInvitesSearch_happy(t *testing.T) {
	d := &SearchReply{}

	err := json.Unmarshal(jsonSearchReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := InvitesSearchRequest{}

	mux.HandleFunc(fmt.Sprintf("/invites/"),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			testURL(t, r, fmt.Sprintf("/invites/"))
			w.Write(jsonSearchReply)
		},
	)

	reply, _, err := client.Invites.Search(context.TODO(), req)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestInvitesSearch_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := InvitesSearchRequest{}

	mux.HandleFunc(fmt.Sprintf("/invites/"),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "POST")
			testURL(t, r, fmt.Sprintf("/invites/"))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Invites.Search(context.TODO(), req)
	assert.Equal(t, err.Error(), d.Error())
}
