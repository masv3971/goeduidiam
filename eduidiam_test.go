package goeduidiam

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mockServer(t *testing.T, mux *http.ServeMux) *httptest.Server {
	return httptest.NewServer(mux)
}

func mockSetup(t *testing.T) (*http.ServeMux, *httptest.Server, *Client) {
	mux := http.NewServeMux()

	//server := httptest.NewServer(mux)
	server := mockServer(t, mux)

	client := mockNew(t, server.URL)

	return mux, server, client
}

func mockNew(t *testing.T, url string) *Client {
	cfg := Config{
		URL: url,
	}
	client := New(cfg)
	return client
}

func testMethod(t *testing.T, r *http.Request, want string) {
	assert.Equal(t, want, r.Method)
}

func testURL(t *testing.T, r *http.Request, want string) {
	assert.Equal(t, want, r.RequestURI)
}

func testBody(t *testing.T, r *http.Request, want string) {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	assert.NoError(t, err)

	got := buffer.String()
	require.JSONEq(t, want, got)
}
