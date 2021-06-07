package http

import (
	"context"
	"fmt"
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	"log"
)

func (p *Provider) Setup() {
	// TODO: Setup the HTTP server
	log.Println(fmt.Sprintf("Starting HTTP Server on port=%s", p.port))

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	p.server = engine

	// todo: middleware
	// todo: logger middleware
	// todo: limit https://github.com/aviddiviner/gin-limit (add default? 1,000?)

	// register all the endpoints with the server
	for _, e := range p.endpoints {
		endpoint := e // loop variable capture issue if we dont copy e into endpoint

		log.Println(fmt.Sprintf("Registering [%s] %s to HTTP Server", endpoint.method, endpoint.path))
		p.server.Handle(endpoint.method, endpoint.path, func(ctx *gin.Context) {
			endpoint.handler(NewContext(ctx))
		})
	}
}

func (p *Provider) Run(ctx context.Context) {
	// todo: handle error here? ?(panic?)
	_ = manners.ListenAndServe(fmt.Sprintf(":%s", p.port), p.server)
}

func (p *Provider) Shutdown() {
	// manners doesn't wait for requests to finish so it will kill any that are running
	manners.Close()
	p.finished = true
	log.Println("Finished shutting down HTTP Server")
}

func (p *Provider) IsFinished() bool {
	return p.finished
}
