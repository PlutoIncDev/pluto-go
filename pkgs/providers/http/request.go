package http

import (
	"io"
	"net/http"
)

func NewRequest(e *Endpoint, body io.Reader) (*http.Request, error) {
	return http.NewRequest(e.Method, e.Path, body)
}
