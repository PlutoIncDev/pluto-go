package httptest

import (
	"context"
	"io"
	nethttp "net/http"
	//nethttptest "net/http/httptest"
	"pluto/pkgs/providers/http"
)

type Request struct {
	*nethttp.Request
}

func NewRequest(method http.Method, path string, body io.Reader) (*Request, error) {
	req, err := nethttp.NewRequest(string(method), path, body)
	return &Request{req}, err
}

func NewRequestWithContext(ctx context.Context, method http.Method, path string, body io.Reader) (*Request, error) {
	req, err := nethttp.NewRequestWithContext(ctx, string(method), path, body)

	return &Request{req}, err
}