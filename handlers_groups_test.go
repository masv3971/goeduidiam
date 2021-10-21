package goeduidiam

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupsGetAll_happy(t *testing.T) {
	d := &GroupsReply{}

	err := json.Unmarshal(jsonGroupsReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	mux.HandleFunc(fmt.Sprintf("/groups/"),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/groups/"))
			w.Write(jsonGroupsReply)
		},
	)

	reply, _, err := client.Groups.GetAll(context.TODO())
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestGroupsGetAll_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	mux.HandleFunc(fmt.Sprintf("/groups/"),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/groups/"))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Groups.GetAll(context.TODO())
	assert.Equal(t, err.Error(), d.Error())
}

func TestGroupsGet_happy(t *testing.T) {
	d := &GroupsReply{}

	err := json.Unmarshal(jsonGroupsReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := GroupsGetRequest{
		ScimID: "test",
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s", req.ScimID),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/groups/%s", req.ScimID))
			w.Write(jsonGroupsReply)
		},
	)

	reply, _, err := client.Groups.Get(context.TODO(), req)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestGroupsGet_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := GroupsGetRequest{
		ScimID: "test",
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s", req.ScimID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/groups/%s", req.ScimID))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Groups.Get(context.TODO(), req)
	assert.Equal(t, err.Error(), d.Error())
}

func TestGroupsPost_happy(t *testing.T) {
	d := &GroupsReply{}

	err := json.Unmarshal(jsonGroupsReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := GroupsPostRequest{}

	mux.HandleFunc(fmt.Sprintf("/groups/"),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			testURL(t, r, fmt.Sprintf("/groups/"))
			w.Write(jsonGroupsReply)
		},
	)

	reply, _, err := client.Groups.Post(context.TODO(), req)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestGroupsPost_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := GroupsPostRequest{}

	mux.HandleFunc(fmt.Sprintf("/groups/"),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "POST")
			testURL(t, r, fmt.Sprintf("/groups/"))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Groups.Post(context.TODO(), req)
	assert.Equal(t, err.Error(), d.Error())
}

func TestGroupsPut_happy(t *testing.T) {
	d := &GroupsReply{}

	err := json.Unmarshal(jsonGroupsReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := GroupsPutRequest{
		ScimID: "test",
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s", req.ScimID),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "PUT")
			testURL(t, r, fmt.Sprintf("/groups/%s", req.ScimID))
			w.Write(jsonGroupsReply)
		},
	)

	reply, _, err := client.Groups.Put(context.TODO(), req)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestGroupsPut_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := GroupsPutRequest{
		ScimID: "test",
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s", req.ScimID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "PUT")
			testURL(t, r, fmt.Sprintf("/groups/%s", req.ScimID))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Groups.Put(context.TODO(), req)
	assert.Equal(t, err.Error(), d.Error())
}

func TestGroupsDelete_happy(t *testing.T) {
	d := &EmptyStruct{}

	err := json.Unmarshal(jsonEmpty, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := GroupsDeleteRequest{
		ScimID: "test",
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s", req.ScimID),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
			testURL(t, r, fmt.Sprintf("/groups/%s", req.ScimID))
			w.Write(jsonEmpty)
		},
	)

	reply, _, err := client.Groups.Delete(context.TODO(), req)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestGroupsDelete_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := GroupsDeleteRequest{
		ScimID: "test",
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s", req.ScimID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "DELETE")
			testURL(t, r, fmt.Sprintf("/groups/%s", req.ScimID))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Groups.Delete(context.TODO(), req)
	assert.Equal(t, err.Error(), d.Error())
}

func TestGroupsSearch_happy(t *testing.T) {
	d := &SearchReply{}

	err := json.Unmarshal(jsonSearchReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := GroupsSearchRequest{}

	mux.HandleFunc(fmt.Sprintf("/groups/"),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			testURL(t, r, fmt.Sprintf("/groups/"))
			w.Write(jsonSearchReply)
		},
	)

	reply, _, err := client.Groups.Search(context.TODO(), req)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestGroupsSearch_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := GroupsSearchRequest{}

	mux.HandleFunc(fmt.Sprintf("/groups/"),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "POST")
			testURL(t, r, fmt.Sprintf("/groups/"))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Groups.Search(context.TODO(), req)
	assert.Equal(t, err.Error(), d.Error())
}
