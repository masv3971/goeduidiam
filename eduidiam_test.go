package goeduidiam

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mockGenericEndpointServer(t *testing.T, mux *http.ServeMux, method, url string, payload []byte, statusCode int) {
	mux.HandleFunc(fmt.Sprintf(url),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(statusCode)
			testMethod(t, r, method)
			testURL(t, r, url)
			w.Write(payload)
		},
	)
}

func mockServer(t *testing.T, mux *http.ServeMux) *httptest.Server {
	return httptest.NewServer(mux)
}

func mockSetup(t *testing.T) (*http.ServeMux, *httptest.Server, *Client) {
	mux := http.NewServeMux()

	server := mockServer(t, mux)

	client := mockNew(t, server.URL)

	return mux, server, client
}

func mockNew(t *testing.T, url string) *Client {
	cfg := Config{
		URL: url,
		Token: TokenConfig{
			Certificate: []byte{},
			PrivateKey:  []byte{},
			Password:    "testPassword",
			Scope:       "testScope",
			Type:        "testType",
			URL:         url,
			Key:         "testKey",
			Client:      "testClient",
		},
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
