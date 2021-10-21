package goeduidiam

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusGetHealthy_happy(t *testing.T) {
	d := &HealthyReply{}

	err := json.Unmarshal(jsonHealthyReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	mux.HandleFunc(fmt.Sprintf("/status/healthy/"),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/status/healthy/"))
			w.Write(jsonHealthyReply)
		},
	)

	reply, _, err := client.Status.GetHealthy(context.TODO())
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestStatusGetHealthy_sad(t *testing.T) {
	d := &Errors{}

	err := json.Unmarshal(jsonErrorReply, d)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	mux, server, client := mockSetup(t)
	defer server.Close()

	mux.HandleFunc(fmt.Sprintf("/status/healthy/"),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/status/healthy/"))
			w.Write(jsonErrorReply)
		},
	)

	_, _, err = client.Status.GetHealthy(context.TODO())
	assert.Equal(t, err.Error(), d.Error())
}
