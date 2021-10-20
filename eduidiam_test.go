package goeduidiam

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockSetup(t *testing.T) (*http.ServeMux, *httptest.Server, *Client) {
	mux := http.NewServeMux()

	//	server := httptest.NewTLSServer(mux)
	server := httptest.NewServer(mux)

	client := mockNew(t, server.URL)

	return mux, server, client
}

func mockNew(t *testing.T, url string) *Client {
	cfg := Config{
		URL: url,
	}
	client, err := New(cfg)
	assert.NoError(t, err)
	return client
}
