package http

import (
	"github.com/gin-gonic/gin"
)

type endpointHandler func(ctx *Context)

type endpoint struct {
	method  string
	path    string
	handler endpointHandler
}

type Provider struct {
	port     string
	finished bool

	endpoints []*endpoint
	server    *gin.Engine
}

type Context struct {
	HTTP *gin.Context
}

type H gin.H
