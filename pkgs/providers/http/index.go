package http

import (
	"fmt"
	"log"
	"time"
)

type Provider struct {
	port string
	endpoints []*endpoint
}

func NewProvider(port string) *Provider {
	return &Provider{
		port: port,
	}
}

func (p *Provider) Setup() {
	log.Println("HTTP Provider is Setting Up")

	// TODO: Setup the HTTP server
}

func (p *Provider) Run() {
	log.Println(fmt.Sprintf("HTTP Server is Running on localhost:%s", p.port))

	// TODO: Run the HTTP server
	for i := 0; i < 6; i++ {
		time.Sleep(time.Second)
		log.Println("Im running!")
	}
}

func (p *Provider) Shutdown() {
	log.Println("HTTP Provider is Shutting Down")

	// TODO: Shutdown the HTTP Server
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

// TODO: add middleware?