package http

import (
	"context"
	"fmt"
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (p *Provider) Setup() {
	// TODO: Setup the HTTP server
	log.Info(fmt.Sprintf("Setting up HTTP Server on port=%s with mode=%s", p.port, p.ginMode))

	gin.SetMode(p.ginMode)

	engine := gin.New()
	p.server = engine

	// todo: limit https://github.com/aviddiviner/gin-limit (add default? 1,000?)
	p.setupMiddlewares()
	p.setupEndpoints()
}

func (p *Provider) Run(ctx context.Context) {
	// todo: handle error here? ?(panic?)
	log.Info("Running HTTP Server")
	if p.ginMode != gin.TestMode {
		_ = manners.ListenAndServe(fmt.Sprintf(":%s", p.port), p.server)
	}
}

func (p *Provider) Shutdown() {
	// manners doesn't wait for requests to finish so it will kill any that are running
	if p.ginMode != gin.TestMode {
		manners.Close()
	}
	p.finished = true
	log.Info("Finished shutting down HTTP Server")
}

func (p *Provider) IsFinished() bool {
	return p.finished
}

func (p *Provider) setupMiddlewares() {
	for _, m := range p.middlewares {
		middleware := m // loop variable capture issue
		p.server.Use(func(c *gin.Context) {
			middleware(NewContext(c))
		})
	}
}

func (p *Provider) setupEndpoints() {
	// register all the endpoints with the server
	for _, e := range p.endpoints {
		endpoint := e // loop variable capture issue

		log.Info(fmt.Sprintf("Registering [%s] %s to HTTP Server", endpoint.Method, endpoint.Path))

		p.server.Handle(endpoint.Method, endpoint.Path, func(ctx *gin.Context) {
			// THIS FUNCTION WILL ONLY FIRE ON REGISTERED ROUTES,
			// e.g. if a route is not defined, this function will not be called
			// register a middleware instead!
			endpoint.Handler(NewContext(ctx))
		})
	}
}
