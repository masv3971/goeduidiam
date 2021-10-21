package goeduidiam

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersGet_happy(t *testing.T) {
	d := &UsersReply{}

	err := json.Unmarshal(jsonUsersReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := UsersGetRequest{
		ScimID: "testScimID",
	}

	mux.HandleFunc(fmt.Sprintf("/users/%s", req.ScimID),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/users/%s", req.ScimID))
			w.Write(jsonUsersReply)
		},
	)

	reply, _, err := client.Users.Get(context.TODO(), req)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestUsersGet_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := UsersGetRequest{
		ScimID: "testScimID",
	}

	mux.HandleFunc(fmt.Sprintf("/users/%s", req.ScimID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/users/%s", req.ScimID))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Users.Get(context.TODO(), req)
	assert.Equal(t, err.Error(), d.Error())
}

func TestUsersPut_happy(t *testing.T) {
	d := &UsersReply{}

	err := json.Unmarshal(jsonUsersReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := UsersPutRequest{
		ScimID: "testScimID",
	}

	mux.HandleFunc(fmt.Sprintf("/users/%s", req.ScimID),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "PUT")
			testURL(t, r, fmt.Sprintf("/users/%s", req.ScimID))
			w.Write(jsonUsersReply)
		},
	)

	reply, _, err := client.Users.Put(context.TODO(), req)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}
func TestUsersPut_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := UsersPutRequest{
		ScimID: "testID",
		Data:   UsersRequest{},
	}

	mux.HandleFunc(fmt.Sprintf("/users/%s", req.ScimID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "PUT")
			testURL(t, r, fmt.Sprintf("/users/%s", req.ScimID))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Users.Put(context.TODO(), req)
	assert.Equal(t, err.Error(), d.Error())
}
func TestUsersPost_happy(t *testing.T) {
	d := &UsersReply{}

	err := json.Unmarshal(jsonUsersReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := UsersPostRequest{}

	mux.HandleFunc(fmt.Sprintf("/users/"),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			testURL(t, r, fmt.Sprintf("/users/"))
			w.Write(jsonUsersReply)
		},
	)

	reply, _, err := client.Users.Post(context.TODO(), req)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestUsersPost_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := UsersPostRequest{}

	mux.HandleFunc(fmt.Sprintf("/users/"),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "POST")
			testURL(t, r, fmt.Sprintf("/users/"))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Users.Post(context.TODO(), req)
	assert.Equal(t, err.Error(), d.Error())
}

func TestUsersSearch_happy(t *testing.T) {
	d := &UsersReply{}

	err := json.Unmarshal(jsonUsersReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := UsersSearchRequest{}

	mux.HandleFunc(fmt.Sprintf("/users/"),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			testURL(t, r, fmt.Sprintf("/users/"))
			w.Write(jsonUsersReply)
		},
	)

	reply, _, err := client.Users.Search(context.TODO(), req)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestUsersSearch_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := UsersSearchRequest{}

	mux.HandleFunc(fmt.Sprintf("/users/"),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "POST")
			testURL(t, r, fmt.Sprintf("/users/"))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Users.Search(context.TODO(), req)
	assert.Equal(t, err.Error(), d.Error())
}
