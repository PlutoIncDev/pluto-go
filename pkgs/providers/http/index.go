package http

import (
	"fmt"
	"net/http"
	"pluto/pkgs/internal/httpserver"
)

func NewProvider(port string) *Provider {
	return &Provider{
		port: port,
	}
}

func (p *Provider) Setup() {
	// TODO: Setup the HTTP server
	p.server = httpserver.NewHTTPServer()

	// register all the endpoints with the server
	for _, e := range p.endpoints {
		p.server.RegisterHandlerFunc(string(e.method), e.path, func(writer http.ResponseWriter, request *http.Request) {
			e.handler(NewContext())
		})
	}
}

func (p *Provider) Run() {
	// todo: Run HTTP Server
	_ = p.server.Start()
}

func (p *Provider) Shutdown() {
	// TODO: Shutdown the HTTP Server
}

func (p *Provider) RegisterEndpoint(method Method, path string, handler endpointHandler) {
	// check if the endpoint is already registered
	for _, e := range p.endpoints {
		if e.path == path && e.method == method {
			panic(fmt.Sprintf("Endpoint with method %s and path %s should only be registered once", method, path))
		}
	}

	e := NewEndpoint(method, path, handler)
	p.endpoints = append(p.endpoints, e)
}

func (p *Provider) GetEndpoints() []*endpoint {
	return p.endpoints
}

// TODO: add middleware?
