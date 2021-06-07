package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func NewProvider(port string) *Provider {
	return &Provider{
		port: port,
	}
}

func (p *Provider) Setup() {
	// TODO: Setup the HTTP server
	log.Println(fmt.Sprintf("Starting HTTP Server on port=%s", p.port))

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	p.server = engine

	// todo: middleware
	// todo: logger middleware

	// register all the endpoints with the server
	for _, e := range p.endpoints {
		// TODO: test that this loop variable capture issue is fixed!
		endpoint := e // loop variable capture issue if we dont copy e into endpoint

		log.Println(fmt.Sprintf("Registering [%s] %s to HTTP Provider", endpoint.method, endpoint.path))
		p.server.Handle(string(endpoint.method), endpoint.path, func(ctx *gin.Context) {
			endpoint.handler(NewContext(ctx))
		})
	}
}

func (p *Provider) Run() {
	// todo: Run HTTP Server
	_ = p.server.Run(fmt.Sprintf(":%s", p.port))
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
	log.Println(p.endpoints)
}

func (p *Provider) GetEndpoints() []*endpoint {
	return p.endpoints
}

// TODO: add middleware?
