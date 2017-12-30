package dingbot

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type RewriteTransport struct {
	Transport http.RoundTripper
}

func (t *RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	if t.Transport == nil {
		return http.DefaultTransport.RoundTrip(req)
	}
	return t.Transport.RoundTrip(req)
}

func testServer() (*http.Client, *http.ServeMux, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	transport := &RewriteTransport{&http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}}

	client := &http.Client{Transport: transport}
	return client, mux, server
}

func assertMethod(t *testing.T, expectedMethod string, req *http.Request) {
	assert.Equal(t, expectedMethod, req.Method)
}

func assertAPIKey(t *testing.T, req *http.Request) {
	queryValues := req.URL.Query()
	_, ok := queryValues["access_token"]
	assert.Contains(t, true, ok)
}
