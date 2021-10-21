package goeduidiam

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginPost_happy(t *testing.T) {
	d := &EmptyStruct{}

	err := json.Unmarshal(jsonEmpty, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := LoginPostRequest{
		Data: LoginRequest{
			DataOwner: "test",
		},
	}

	mux.HandleFunc(fmt.Sprintf("/login/"),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			testURL(t, r, fmt.Sprintf("/login/"))
			w.Write(jsonEmpty)
		},
	)

	reply, _, err := client.Login.Post(context.TODO(), req)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestLoginPost_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	req := LoginPostRequest{
		Data: LoginRequest{
			DataOwner: "testDataOwner",
		},
	}

	mux.HandleFunc(fmt.Sprintf("/login/"),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "POST")
			testURL(t, r, fmt.Sprintf("/login/"))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Login.Post(context.TODO(), req)
	assert.Equal(t, err.Error(), d.Error())
}
