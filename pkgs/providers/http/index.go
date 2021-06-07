package http

import (
	"fmt"
	"log"
	"pluto/pkgs/internal/http"
)

type Provider struct {
	port string
	server *http.Server
	endpoints []*endpoint
}

func NewProvider(port string) *Provider {
	return &Provider{
		port: port,
	}
}

func (p *Provider) Setup() {
	log.Println("HTTP Provider is Setting Up")

	// Setup the HTTP server
	p.server = http.NewServer()
}

func (p *Provider) Run() {
	log.Println(fmt.Sprintf("HTTP Server is Running on localhost:%s", p.port))
}

func (p *Provider) Shutdown() {
	log.Println("HTTP Provider is Shutting Down")
}

func (p *Provider) RegisterEndpoint(method Method, path string, handler func()) {
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