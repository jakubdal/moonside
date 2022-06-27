package internal

import (
	"io"
	"io/ioutil"
	"net/http"
)

// HTTPDoer abstracts the Do method of http.Client struct.
//
// It allows easy mocking of http client in tests.
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// CleanupHTTPResponse discards http response body to reuse TCP connection.
//
// Source: answer to https://github.com/golang/go/issues/48860
func CleanupHTTPResponse(resp *http.Response) {
	// ignore error - we cannot handle it anyways
	io.Copy(ioutil.Discard, resp.Body)
	// ignore error again - we still cannot handle it
	resp.Body.Close()
}
