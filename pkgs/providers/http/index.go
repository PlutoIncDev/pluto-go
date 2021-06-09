package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func NewProvider(port string) *Provider {
	p := &Provider{
		// config
		port: port,

		// internal
		finished: false,
		ginMode:  gin.ReleaseMode,
	}

	return p
}

func (p *Provider) SetTestMode() {
	p.ginMode = gin.TestMode
}

func (p *Provider) RegisterEndpoint(e *Endpoint) {
	p.endpoints = append(p.endpoints, e)
}

func (p *Provider) RegisterMiddleware(middleware Handler) {
	p.middlewares = append(p.middlewares, middleware)
}

func (p *Provider) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	p.server.ServeHTTP(w, req)
}

func (p *Provider) NewRecorder() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}
